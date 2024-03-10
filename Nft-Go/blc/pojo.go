package blc

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	"math/big"
)

func init() {
	hessian.RegisterPOJO(&UserKey{})
	hessian.RegisterPOJO(&GiveDTO{})
	hessian.RegisterPOJO(&GetDcFromActivityDTO{})
	hessian.RegisterPOJO(&DcHistoryAndMessageOutputDTO{})
	hessian.RegisterPOJO(&CreatePoolDTO{})
	hessian.RegisterPOJO(&CreateActivityDTO{})
	hessian.RegisterPOJO(&CheckDcAndReturnTimeOutputDTO{})
	hessian.RegisterPOJO(&CheckDcAndReturnTimeDTO{})
	hessian.RegisterPOJO(&BeforeMintDTO{})
	hessian.RegisterPOJO(&ActivityAndPool{})
	hessian.RegisterPOJO(&Pool{})
	hessian.RegisterPOJO(&Dc{})
	hessian.RegisterPOJO(&Activity{})
	hessian.RegisterPOJO(&TraceStruct{})

}

type UserKey struct {
	UserKey string
}

type GiveDTO struct {
	ToAddress string
	DcId      int
}
type GetDcFromActivityDTO struct {
	ActivityId *big.Int
	Password   []byte
}

type DcHistoryAndMessageOutputDTO struct {
	Args           []TraceStruct
	Hash           []byte
	CreatorAddress string
	OwnerAddress   string
	DcName         string
	PoolId         *big.Int
}

type CreatePoolDTO struct {
	LimitAmount *big.Int
	Price       *big.Int
	Amount      *big.Int
	Cid         string
	DcName      string
}

type CreateActivityDTO struct {
	Name     string
	Password []byte
	Amount   *big.Int
	Cid      string
	DcName   string
}

type CheckDcAndReturnTimeOutputDTO struct {
	CheckResult bool
	TimeList    []*big.Int
}

type CheckDcAndReturnTimeDTO struct {
	Owner          string
	CollectionHash [][]byte
}

type BeforeMintDTO struct {
	DcId     *big.Int
	UniqueId []byte
}

type ActivityAndPool struct {
	Activity *Activity
	Pool     *Pool
}

type Pool struct {
	Cid         string
	Name        string
	Price       *big.Int
	Amount      *big.Int
	Left        *big.Int
	LimitAmount *big.Int
	Creator     string
	CreateTime  *big.Int
}

type Dc struct {
	UniqueHash  []byte
	MintTime    *big.Int
	PoolId      *big.Int
	IndexInPool *big.Int
}

type Activity struct {
	Name       string
	EncodedKey []byte
	PoolId     *big.Int
}

type TraceStruct struct {
	Sender        string
	To            string
	OperateTime   *big.Int
	OperateRecord string
}

func (k *UserKey) JavaClassName() string {
	return "com.topview.entity.dto.UserKey"
}
func (k GiveDTO) JavaClassName() string {
	return "com.topview.entity.dto.GiveDTO"
}
func (k GetDcFromActivityDTO) JavaClassName() string {
	return "com.topview.entity.dto.GetDcFromActivityDTO"
}
func (k DcHistoryAndMessageOutputDTO) JavaClassName() string {
	return "com.topview.entity.dto.DcHistoryAndMessageOutputDTO"
}
func (k CreatePoolDTO) JavaClassName() string {
	return "com.topview.entity.dto.CreatePoolDTO"
}
func (k CreateActivityDTO) JavaClassName() string {
	return "com.topview.entity.dto.CreateActivityDTO"
}
func (k CheckDcAndReturnTimeOutputDTO) JavaClassName() string {
	return "com.topview.entity.dto.CheckDcAndReturnTimeOutputDTO"
}
func (k CheckDcAndReturnTimeDTO) JavaClassName() string {
	return "com.topview.entity.dto.CheckDcAndReturnTimeDTO"
}
func (k BeforeMintDTO) JavaClassName() string {
	return "com.topview.entity.dto.BeforeMintDTO"
}
func (k ActivityAndPool) JavaClassName() string {
	return "com.topview.entity.dto.ActivityAndPool"
}
func (k Pool) JavaClassName() string {
	return "com.topview.entity.dto.Pool"
}
func (k Dc) JavaClassName() string {
	return "com.topview.entity.dto.Dc"
}
func (k Activity) JavaClassName() string {
	return "com.topview.entity.dto.Activity"
}
func (k TraceStruct) JavaClassName() string {
	return "com.topview.entity.dto.TraceStruct"
}
