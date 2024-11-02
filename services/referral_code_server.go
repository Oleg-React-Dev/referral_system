package services

import (
	"referral_app/domain/referral_codes"
	"referral_app/domain/users"
	"referral_app/utils/errors"
)

var ReferralService referralCodeInterface = &referralService{}

type referralService struct{}

type referralCodeInterface interface {
	CreateReferralCode(referral_codes.ReferralCode) (*referral_codes.ReferralCode, *errors.RestErr)
	DeleteReferralCode(string) *errors.RestErr
	GetReferralCodeByEmail(string) (*referral_codes.ReferralCode, *errors.RestErr)
	GetReferralsByReferrerId(referrerId string) (users.Users, *errors.RestErr)
}

func (s *referralService) CreateReferralCode(rc referral_codes.ReferralCode) (*referral_codes.ReferralCode, *errors.RestErr) {
	if err := rc.Validate(); err != nil {
		return nil, err
	}

	if err := rc.Save(); err != nil {
		return nil, err
	}
	return &rc, nil
}

func (s *referralService) DeleteReferralCode(userId string) *errors.RestErr {
	rc := referral_codes.ReferralCode{UserId: userId}
	if err := rc.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *referralService) GetReferralCodeByEmail(email string) (*referral_codes.ReferralCode, *errors.RestErr) {
	rc := referral_codes.ReferralCode{}
	if err := rc.GetCodeByEmail(email); err != nil {
		return nil, err
	}
	return &rc, nil
}

func (s *referralService) GetReferralsByReferrerId(referrerId string) (users.Users, *errors.RestErr) {
	var dao referral_codes.ReferralCode
	return dao.FindById(referrerId)
}
