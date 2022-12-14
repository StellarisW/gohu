package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"main/app/common/log"
	"net/http"

	"main/app/service/question/rpc/crud/internal/svc"
	"main/app/service/question/rpc/crud/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteQuestionLogic) DeleteQuestion(in *pb.DeleteQuestionReq) (res *pb.DeleteQuestionRes, err error) {
	logger := log.GetSugaredLogger()
	logger.Debugf("recv message: %v", in.String())

	err = l.svcCtx.Rdb.Del(l.ctx,
		fmt.Sprintf("question_subject_%d", in.QuestionId)).Err()
	if err != nil {
		logger.Errorf("delete [question_subject] cache failed, err: %v", err)
		res = &pb.DeleteQuestionRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	}
	l.svcCtx.Rdb.Del(l.ctx,
		fmt.Sprintf("question_content_%d", in.QuestionId))
	if err != nil {
		logger.Errorf("delete [question_subject] cache failed, err: %v", err)
		res = &pb.DeleteQuestionRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	}

	// 删除question_subject后, 级联删除关联的回答和回答下的评论
	questionSubjectModel := l.svcCtx.QuestionModel.QuestionSubject

	questionSubject, err := questionSubjectModel.WithContext(l.ctx).
		Select(questionSubjectModel.ID, questionSubjectModel.UserID).
		Where(questionSubjectModel.ID.Eq(in.QuestionId)).
		First()
	if err != nil {
		logger.Errorf("query [question_subject] record failed, err: %v", err)
		res = &pb.DeleteQuestionRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	}

	_, err = questionSubjectModel.WithContext(l.ctx).Where(questionSubjectModel.ID.Eq(in.QuestionId)).Delete()
	switch err {
	case gorm.ErrRecordNotFound:
		res = &pb.DeleteQuestionRes{
			Code: http.StatusBadRequest,
			Msg:  "question not found",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	case nil:

	default:
		logger.Errorf("update question failed, err: %v", err)
		res = &pb.DeleteQuestionRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	}

	err = l.svcCtx.Rdb.SRem(l.ctx,
		fmt.Sprintf("question_id_user_set_%v", questionSubject.UserID),
		in.QuestionId).Err()
	if err != nil {
		logger.Errorf("update [question_id_user_set] failed, err: %v", err)
		res = &pb.DeleteQuestionRes{
			Code: http.StatusInternalServerError,
			Msg:  "internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	}

	res = &pb.DeleteQuestionRes{
		Code: http.StatusOK,
		Msg:  "delete question successfully",
		Ok:   true,
	}
	logger.Debugf("send message: %v", res.String())
	return res, nil
}
