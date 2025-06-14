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

		return 0, 0, fmt.Errorf("Больше 2 жлементов в слайсе")
	}
	steps, err := strconv.Atoi(sl[0])
	if err != nil || steps <= 0 {
		log.Println(err)
		return 0, 0, fmt.Errorf("Неудалось преобразовать строку в число или количество шагов меньше или равно 0 ")

	}

	minutHour, err := time.ParseDuration(sl[1])
	if err != nil || minutHour <= 0 {
		log.Println(err)
		return 0, 0, fmt.Errorf("Неудалось преобразовать время в time.Duration или время равно 0")
	}

	return steps, minutHour, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, minutHour, _ := parsePackage(data)
	distance := (float64(steps) * stepLength) / mInKm
	calories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, minutHour)
	if steps <= 0 || weight <= 0 || height <= 0 || calories <= 0 || distance <= 0 {
		err := errors.New("one of the parametrs is 0")
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
