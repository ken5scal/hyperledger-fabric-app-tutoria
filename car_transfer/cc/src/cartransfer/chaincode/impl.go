package chaincode

import (
	"cartransfer"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

type CarTransferCC struct {
}

func (this *CarTransferCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("cartransfer")
	logger.Info(("chaincode initialized"))
	return shim.Success([]byte{})
}

func (this *CarTransferCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("cartransfer")
	timeStamp, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get TX timestamp: %s", err))
	}

	logger.Infof(
		"Invoke called: Tx ID = %s, timestamp = %s",
		stub.GetTxID(),
		timeStamp)

	var (
		fcn  string
		args []string
	)

	fcn, args = stub.GetFunctionAndParameters()
	logger.Infof("function name = %s", fcn)

	switch fcn {
	case "AddCar":
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		owner := new(cartransfer.Owner)
		err := json.Unmarshal([]byte(args[0]), owner)
		if err != nil {
			mes := fmt.Sprintf("failed to un,arshal Owner JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.AddOwner(stub, owner)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte{})
	case "ListCars":
		cars, err := this.ListCars(stub)
		if err != nil {
			return shim.Error(err.Error())
		}

		b, err := json.Marshal(cars)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal Cars: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		return shim.Success(b)

	case "GetCar":
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		var id string
		err := json.Unmarshal([]byte(args[0]), &id)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 1st arg: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		car, err := this.GetCar(stub, id)
		if err != nil {
			return shim.Error(err.Error())
		}

		b, err := json.Marshal(car)
		if err != nil {
			mes := fmt.Sprintf("failed to marshal Car: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		return shim.Success(b)

	case "UpdateCar":
		if err := checkLen(logger, 1, args); err != nil {
			return shim.Error(err.Error())
		}

		car := new(cartransfer.Car)
		err := json.Unmarshal([]byte(args[0]), car)
		if err != nil {
			mes := fmt.Sprintf("failed to un,arshal Car JSON: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.UpdateCar(stub, car)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte{})

	case "TransferCar":
		if err := checkLen(logger, 2, args); err != nil {
			return shim.Error(err.Error())
		}

		var carId, newOwnerId string
		err := json.Unmarshal([]byte(args[0]), &carId)
		if err != nil {
			mes := fmt.Sprintf(
				"failed to unmarshal the 1st arrgument: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = json.Unmarshal([]byte(args[1]), &newOwnerId)
		if err != nil {
			mes := fmt.Sprintf("failed to unmarshal the 2nd arg: %s", err.Error())
			logger.Warning(mes)
			return shim.Error(mes)
		}

		err = this.TransferCar(stub, carId, newOwnerId)
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte{})
	}

	mes := fmt.Sprintf("Unknown method: %s", fcn)
	logger.Warning(mes)
	return shim.Error(mes)
}

func checkLen(logger *shim.ChaincodeLogger, expected int, args []string) error {
	if len(args) < expected {
		mes := fmt.Sprintf(
			"not enough number of arguments: %d given, %d expected",
			len(args),
			expected,
		)
		logger.Warning(mes)
		return errors.New(mes)
	}
	return nil
}

// Gets CN of transaction creator
func getCreator(stub shim.ChaincodeStubInterface) (string, error) {
	logger := shim.NewLogger("cartransfer")
	creator, err := stub.GetCreator()
	if err != nil {
		mes := fmt.Sprintf("failed to get Creator :%s", err)
		logger.Error(mes)
		return "", errors.New(mes)
	}

	in := &msp.SerializedIdentity{}
	err = proto.Unmarshal(creator, in)
	if err != nil {
		mes := fmt.Sprintf("failed to unmarshal creator: %s", err)
		logger.Error(mes)
		return "", errors.New(mes)
	}
	pemBytes := in.IdBytes
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		mes := "failed to decode creator pem"
		logger.Error(mes)
		return "", errors.New(mes)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		mes := fmt.Sprintf("failed to decode x509 certificate: %s", err)
		logger.Error(mes)
		return "", errors.New(mes)
	}

	return cert.Subject.CommonName, nil
}

func (this *CarTransferCC) AddOwner(stub shim.ChaincodeStubInterface, owner *cartransfer.Owner) error {
	return errors.New("not implemented yed")
}

func (this *CarTransferCC) CheckOwner(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	return false, errors.New("not implemented yet")
}

func (this *CarTransferCC) ListOwners(stub shim.ChaincodeStubInterface) ([]*cartransfer.Owner, error) {
	return nil, errors.New("not implemented yed")
}

func (this *CarTransferCC) AddCar(stub shim.ChaincodeStubInterface, car *cartransfer.Car) error {
	return errors.New("not implemented yed")
}

func (this *CarTransferCC) CheckCar(stub shim.ChaincodeStubInterface, id string) (bool, error) {
	return false, errors.New("not implemented yed")
}

func (this *CarTransferCC) ValidateCar(stub shim.ChaincodeStubInterface, car *cartransfer.Car) (bool, error) {
	return false, errors.New("not implemented yed")
}

func (this *CarTransferCC) GetCar(stub shim.ChaincodeStubInterface, id string) (*cartransfer.Car, error) {
	return nil, errors.New("not implemented yed")
}

func (this *CarTransferCC) UpdateCar(stub shim.ChaincodeStubInterface, car *cartransfer.Car) error {
	return errors.New("not implemented yed")
}

func (this *CarTransferCC) ListCars(stub shim.ChaincodeStubInterface) ([]*cartransfer.Car, error) {
	return nil, errors.New("not implemented yed")
}

func (this *CarTransferCC) TransferCar(stub shim.ChaincodeStubInterface, carId string, newOwnerId string) error {
	return errors.New("not implemented yed")
}
