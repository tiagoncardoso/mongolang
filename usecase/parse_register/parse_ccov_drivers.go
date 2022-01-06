package parse_register

import (
	"ccovdata/domain/entity/ccov"
	"ccovdata/domain/entity/valida"
	"ccovdata/domain/entity/valida/locker_register"
	"ccovdata/domain/repository"
	"gopkg.in/gookit/color.v1"
)

type ParseRegister struct {
	Repository *repository.ValidaDatabaseRepository
	Register   *ccov.DriverRegister
}

func NewParseRegister(repository *repository.ValidaDatabaseRepository, register *ccov.DriverRegister) *ParseRegister {
	return &ParseRegister{
		Repository: repository,
		Register:   register,
	}
}

func (pr *ParseRegister) SaveDriver() (int64, error) {
	r := pr.Register
	lockerRegister := pr.buildDriverLocker()

	driver, err := valida.NewDriverRegister(
		r.Name,
		r.Document,
		lockerRegister,
		r.State,
		r.CreationTime,
	)

	driver.SetTipoVinculo(r.DriverProfile)

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

	var vm *valida.VehicleRegister
	var carr1 *valida.VehicleRegister
	var carr2 *valida.VehicleRegister
	var carr3 *valida.VehicleRegister

	vm, _ = pr.buildVehicle(0)
	carr1, _ = pr.buildVehicle(1)
	carr2, _ = pr.buildVehicle(2)
	carr3, _ = pr.buildVehicle(3)

	var vid int64
	if vm != nil {
		vm.SetVehicleType(2)
		vid, _ = pr.Repository.InsertVehicle(vm)
		vids = append(vids, vid)
	}
	if carr1 != nil {
		carr1.SetVehicleType(1)
		vid, _ = pr.Repository.InsertVehicle(carr1)
		vids = append(vids, vid)
	}
	if carr2 != nil {
		carr2.SetVehicleType(1)
		vid, _ = pr.Repository.InsertVehicle(carr2)
		vids = append(vids, vid)
	}
	if carr3 != nil {
		carr3.SetVehicleType(1)
		vid, _ = pr.Repository.InsertVehicle(carr3)
		vids = append(vids, vid)
	}

	return vids, nil
}

func (pr *ParseRegister) SaveTravel() (int64, error) {
	r := pr.Register

	travel := valida.NewTravelRegister(r.CreationTime)
	travel.SetValorCarga(10000)
	travel.SetChargerType(r.Product)

	tid, _ := pr.Repository.InsertTravel(travel)

	return tid, nil
}

func (pr *ParseRegister) SaveRegister(driverId int64, vehiclesID []int64, travelId int64) (int64, error) {
	r := pr.Register
	vehiclesSize := len(vehiclesID)

	var vehicle int64
	var carreta1 int64
	var carreta2 int64
	var carreta3 int64

	switch vehiclesSize {
	case 4:
		carreta3 = vehiclesID[3]
		carreta2 = vehiclesID[2]
		carreta1 = vehiclesID[1]
		vehicle = vehiclesID[0]
	case 3:
		carreta2 = vehiclesID[2]
		carreta1 = vehiclesID[1]
		vehicle = vehiclesID[0]
	case 2:
		carreta1 = vehiclesID[1]
		vehicle = vehiclesID[0]
	case 1:
		vehicle = vehiclesID[0]
	}

	reg := valida.NewRegister(driverId, vehicle, carreta1, carreta2, carreta3, travelId, r.RegisterExtra.Protocol, r.CreationTime)
	reg.SetPlus(false)
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

	result := valida.NewResultRegister(registerId, r.CreationTime)
	result.SetSituacao(r.Score)
	result.SetValidade(r.Score, r.CreationTime, r.ValidityTime)

	if result.Situacao == "ADEQUADO" {
		code, _ := pr.Repository.GenerateCode()
		result.SetCode(code)
	}

	rid, err := pr.Repository.InsertResultRegister(result)
	if err != nil {
		return -1, err
	}

	return rid, nil
}

func (pr *ParseRegister) SaveDemand(registerId int64) (int64, error) {
	r := pr.Register

	demand := valida.NewDemandRegister(registerId, r.CompanyPortalId, r.CreationTime)
	rid, err := pr.Repository.InsertDemandRegister(demand)
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
		reg.DateOfBirthday.Format("2006-01-02 15:04:05"),
	)
	lr.SetDocumentsData(
		reg.Document,
		reg.Document3,
		reg.IssueStateDocument3,
		reg.Document2,
		reg.IssueStateDocument2,
		reg.DueDateDocument3.Format("2006-01-02 15:04:05"),
		lr.SetCnhCategory(reg.CategoryDocument3),
	)
	lr.SetContactData(reg.Landline)

	return lr
}

func (pr *ParseRegister) buildVehicle(id int) (*valida.VehicleRegister, error) {
	vs := pr.Register.DeviceRegisters
	qdtVeiculos := len(vs)

	if id >= qdtVeiculos {
		return nil, nil
	}
	lockerRegister := pr.buildVehicleLocker(id)

	veic, _ := valida.NewVehicleRegister(
		vs[id].Plate,
		vs[id].StatePlate,
		lockerRegister,
		pr.Register.CreationTime,
	)
	veic.SetTipoVinculo(pr.Register.DriverProfile)

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
