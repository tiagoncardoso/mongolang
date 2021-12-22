package parse_register

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/entity/valida"
	"ccovdata/domain/entity/valida/locker_register"
	"ccovdata/domain/repository"
	"gopkg.in/gookit/color.v1"
	"time"
)

var NOW = time.Now()

type ParseRegister struct {
	Repository *repository.ValidaDatabaseRepository
	Register   *ccov.DriverRegisterExternal
}

func NewParseRegister(repository *repository.ValidaDatabaseRepository, register *ccov.DriverRegisterExternal) *ParseRegister {
	return &ParseRegister{
		Repository: repository,
		Register:   register,
	}
}

func (pr *ParseRegister) SaveDriver() (int64, error) {
	r := pr.Register
	lockerRegister := pr.buildDriverLocker()

	driver, err := valida.NewDriverRegister(
		r.Driver.Name,
		r.Driver.Document,
		lockerRegister,
		r.Driver.State,
	)

	driver.SetTipoVinculo(r.RegisterExtra.DriverProfile)

	if err != nil {
		return 0, err
	}

	var did int64
	did, err = pr.Repository.InsertDriver(driver)
	if err != nil {
		return 0, err
	}

	return did, nil
}

func (pr *ParseRegister) SaveVehicle() ([]int64, error) {
	var vids []int64

	var vm *valida.VehicleRegister = nil
	var carr1 *valida.VehicleRegister = nil
	var carr2 *valida.VehicleRegister = nil
	var carr3 *valida.VehicleRegister = nil

	vm, _ = pr.buildVehicle(0)
	carr1, _ = pr.buildVehicle(1)
	carr2, _ = pr.buildVehicle(2)
	carr3, _ = pr.buildVehicle(3)

	var vid int64
	if vm != nil {
		vid, _ = pr.Repository.InsertVehicle(vm)
		vids = append(vids, vid)
	}
	if carr1 != nil {
		vid, _ = pr.Repository.InsertVehicle(carr1)
		vids = append(vids, vid)
	}
	if carr2 != nil {
		vid, _ = pr.Repository.InsertVehicle(carr2)
		vids = append(vids, vid)
	}
	if carr3 != nil {
		vid, _ = pr.Repository.InsertVehicle(carr3)
		vids = append(vids, vid)
	}

	return vids, nil
}

func (pr *ParseRegister) SaveTravel() (int64, error) {
	r := pr.Register

	travel := valida.NewTravelRegister()
	travel.SetValorCarga(r.ProductValue)
	travel.SetChargerType(r.Product)

	tid, _ := pr.Repository.InsertTravel(travel)

	return tid, nil
}

func (pr *ParseRegister) SaveRegister(driverId int64, vehiclesID []int64, travelId int64) (int64, error) {
	r := pr.Register

	reg := valida.NewRegister(driverId, vehiclesID[0], vehiclesID[1], vehiclesID[2], vehiclesID[3], travelId)
	reg.SetPlus(r.IsPlus())
	reg.SetRegisterValidity(r.RegisterExtra.CreationTime, r.RegisterExtra.ValidityTime, r.RegisterExtra.Score)
	reg.SetRegisterValidation(r.RegisterExtra.Score)

	rid, err := pr.Repository.InsertRegister(reg)
	if err != nil {
		color.Red.Printf("Erro ao salvar cadastro: %s", err)
	}

	return rid, nil
}

func (pr *ParseRegister) SaveResult(registerId int64) (int64, error) {
	r := pr.Register
	code, _ := pr.Repository.GenerateCode()

	result := valida.NewResultRegister(registerId, code)
	result.SetSituacao(r.RegisterExtra.Score)
	result.SetValidade(r.RegisterExtra.Score, r.CreationTime, r.RegisterExtra.ValidityTime)

	rid, err := pr.Repository.InsertResultRegister(result)
	if err != nil {
		return -1, err
	}

	return rid, nil
}

func (pr *ParseRegister) buildDriverLocker() *locker_register.DriverLockerRegister {
	reg := pr.Register.RegisterExtra
	lr := locker_register.NewDriverLockerRegister()
	lr.SetDriverCategory(reg.DriverProfile)
	lr.SetPersonalData(
		reg.Name,
		reg.MotherName,
		reg.DateOfBirthday.Format("YYYY-MM-DD"),
	)
	lr.SetDocumentsData(
		reg.Document,
		reg.Document3,
		reg.IssueStateDocument3,
		reg.Document2,
		reg.IssueStateDocument2,
		reg.DueDateDocument3.Format("YYYY-MM-DD"),
		lr.SetCnhCategory(reg.CategoryDocument3),
	)
	lr.SetContactData(reg.Landline)

	return lr
}

func (pr *ParseRegister) buildVehicle(id int) (*valida.VehicleRegister, error) {
	vs := pr.Register.DeviceRegisters
	if id >= (len(vs) - 1) {
		return nil, nil
	}
	lockerRegister := pr.buildVehicleLocker(id)

	veic, _ := valida.NewVehicleRegister(
		vs[id].Plate,
		vs[id].StatePlate,
		lockerRegister,
	)
	veic.SetTipoVinculo(pr.Register.RegisterExtra.DriverProfile)

	return veic, nil
}

func (pr *ParseRegister) buildVehicleLocker(idVeic int) *locker_register.VehicleLockerRegister {
	reg := pr.Register.DeviceRegisters[idVeic]
	lr := locker_register.NewVehicleLockerRegister()
	lr.SetVehicleData(
		reg.Plate,
		reg.Document2,
		reg.StatePlate,
	)
	vehicleType := "CAVALO_MECANICO"

	if idVeic > 0 {
		vehicleType = "CARRETA"
	}
	lr.SetVehicleType(vehicleType)

	return lr
}
