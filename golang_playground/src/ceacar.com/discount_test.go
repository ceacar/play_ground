package main

import "github.com/DATA-DOG/godog"

func thereIsOnSaleBicycleWhichCostsPounds(arg1 string, arg2 int) error {
	//return godog.ErrPending
	return nil
}

func iAddThreeBicyclesIntoTheBasket(arg1 string) error {
	//return godog.ErrPending
	return nil
}

func iShouldGetThirdBicycleForFree() error {
	//return godog.ErrPending
	return nil
}

func theBasketTotalShouldBePounds(arg1 int) error {
	//return godog.ErrPending
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is on sale bicycle "([^"]*)" which costs (\d+) pounds$`, thereIsOnSaleBicycleWhichCostsPounds)
	s.Step(`^I add three "([^"]*)" bicycles into the basket$`, iAddThreeBicyclesIntoTheBasket)
	s.Step(`^I should get third bicycle for free$`, iShouldGetThirdBicycleForFree)
	s.Step(`^the basket total should be (\d+) pounds$`, theBasketTotalShouldBePounds)
}
