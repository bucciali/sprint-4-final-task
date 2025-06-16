package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	sl := strings.Split(data, ",")
	if len(sl) != 2 {

		return 0, 0, errors.New("more than 2 elemets in slice")
	}
	steps, err := strconv.Atoi(sl[0])

	if err != nil {

		return 0, 0, fmt.Errorf("mistake with parsing string: %w", err)

	}
	if steps <= 0 {
		return 0, 0, errors.New("steps less than or equal to zero")

	}

	minutHour, err := time.ParseDuration(sl[1])
	if err != nil {
		log.Println(err)
		return 0, 0, fmt.Errorf("mistake with parsing time: %w", err)
	}
	if minutHour <= 0 {
		return 0, 0, errors.New("time less than or equal to zero")
	}

	return steps, minutHour, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, minutHour, _ := parsePackage(data)
	if steps <= 0 {
		err := errors.New("steps less than or equal to zero")
		log.Println(err)
		return ""
	}
	if weight <= 0 {
		err := errors.New("weight less than or equal to zero")
		log.Println(err)
		return ""
	}
	if height <= 0 {
		err := errors.New("height less than or equal to zero")
		log.Println(err)
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	calories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, minutHour)
	if calories <= 0 {
		err := errors.New("calories less than or equal to zero")
		log.Println(err)
		return ""
	}
	if distance <= 0 {
		err := errors.New("distance less than or equal to zero")
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
