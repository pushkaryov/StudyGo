package main

import "fmt"

const PassStatus string = "pass"
const FailStatus string = "fail"

func main() {
	sliceWithStruct := GenerateCheck()

	for i := 0; i < len(sliceWithStruct); i++ {

		if sliceWithStruct[i].Status == PassStatus {
			fmt.Println(sliceWithStruct[i].ServiceID)
		}
	}
}

type HealthCheck struct {
	ServiceID int
	Status    string
}

func GenerateCheck() []HealthCheck {

	var listCheck []HealthCheck

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			var str = HealthCheck{ServiceID: i, Status: PassStatus}
			listCheck = append(listCheck, str)
		} else {
			var str = HealthCheck{ServiceID: i, Status: FailStatus}
			listCheck = append(listCheck, str)
		}

		//switch {
		//case i%2 == 0:
		//	var str = HealthCheck{ServiceID: i, Status: PassStatus}
		//	listCheck = append(listCheck, str)
		//case i%2 != 0:
		//	var str = HealthCheck{ServiceID: i, Status: FailStatus}
		//	listCheck = append(listCheck, str)
		//}
	}

	return listCheck
}
