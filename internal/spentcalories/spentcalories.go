package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	sl := strings.Split(data, ",")
	// обработка ошибки с длиной
	if len(sl) != 3 {
		return 0, "", 0, errors.New("inccoret slice lengh")
	}
	steps, err := strconv.Atoi(sl[0])
	// обработка ошибки с переводом в инт
	if err != nil {
		return 0, "", 0, fmt.Errorf("mistake with convert int type: %w", err)
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("steps less than or equal to zero")
	}
	typeOfActivity := sl[1]

	timeSpent, err := time.ParseDuration(sl[2])
	//обработка ошибки с переводом в тайм дуратион
	if err != nil {
		return 0, "", 0, fmt.Errorf("mistake with parsing time: %w", err)
	}
	if timeSpent <= 0 {
		return 0, "", 0, errors.New("timeSpent less than or equal to zero")
	}
	return steps, typeOfActivity, timeSpent, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	newHeight := height * stepLengthCoefficient

	distance := float64(steps) * newHeight

	return distance / mInKm

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if steps <= 0 {

		return 0
	}

	if height <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)

	hours := duration.Hours()

	return distance / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeOfActivity, timeSpent, _ := parseTraining(data)
	distance := distance(steps, height)
	speed := meanSpeed(steps, height, timeSpent)

	switch typeOfActivity {
	case "Бег":
		caloriesByRunning, _ := RunningSpentCalories(steps, weight, height, timeSpent)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfActivity, timeSpent.Hours(), distance, speed, caloriesByRunning), nil

	case "Ходьба":
		caloriesByWalk, _ := WalkingSpentCalories(steps, weight, height, timeSpent)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfActivity, timeSpent.Hours(), distance, speed, caloriesByWalk), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки\n")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 {
		return 0, fmt.Errorf("steps less than or equal to zero")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight less than or equal to zero")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height less than or equal to zero")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration less than or equal to zero")
	}

	avgSpeed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	wam := weight * avgSpeed * minutes // промежуточный расчет для облегчения читаемости

	return wam / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps less than or equal to zero")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight less than or equal to zero")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height less than or equal to zero")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration less than or equal to zero")
	}

	avgSpeed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	wam := avgSpeed * weight * minutes

	spentCalories := wam / minInH

	return spentCalories * walkingCaloriesCoefficient, nil

}
