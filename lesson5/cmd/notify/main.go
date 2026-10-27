package main

import (
	"errors"
	"fmt"

	"lesson5/student/practice"
)

func main() {
	var scenario int8 = 5
	// Случаи:
	// 1 - "Успех"
	// 2 - "Пустой текст"
	// 3 - "Нет настройки"
	// 4 - "Сломанная реализация"
	// 5 - все случаи

	if scenario > 5 || scenario < 1 {
		fmt.Println("Допустими сценарии от 1 до 5")
	}

	emailNotifier, pushNotifier := newNotifiers()

	switch scenario {
	case 1:
		showSuccess(emailNotifier, pushNotifier)

	case 2:
		// заполним хоть какую-то историю, чтобы она была, это удобно делается в 1 случае
		showSuccess(emailNotifier, pushNotifier)
		allNotifiers := []practice.Notifier{
			emailNotifier,
			pushNotifier,
		}

		err := practice.Dispatch(allNotifiers, "Заказ готов")
		if err != nil {
			fmt.Println("Не успех, хотя должен быть успех, из-за:", err)
			return
		}

		showEmptyMessage(emailNotifier, pushNotifier)

	case 3:
		showMissingConfig()

	case 4:
		showBrokenNotifier()

	case 5:
		showSuccess(emailNotifier, pushNotifier)

		showEmptyMessage(emailNotifier, pushNotifier)
		fmt.Println("---------------------------------------------------------------------------------------------")

		showMissingConfig()

		fmt.Println("---------------------------------------------------------------------------------------------")

		showBrokenNotifier()
	}
}

func newNotifiers() (*practice.EmailNotifier, *practice.PushNotifier) {
	emailNotifier := &practice.EmailNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "Email"},
		Address:      "test@gmail.com",
	}
	pushNotifier := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "Device"},
		DeviceID:     "M1810F7LGD",
	}

	return emailNotifier, pushNotifier
}

func showSuccess(emailNotifier *practice.EmailNotifier, pushNotifier *practice.PushNotifier) {
	allNotifiers := []practice.Notifier{emailNotifier, pushNotifier}
	res := practice.Dispatch(allNotifiers, "Заказ готов")

	if res != nil {
		fmt.Printf("Не успех, хотя должен быть успех, из-за: %v\n", res)
		return
	}

	fmt.Println("Успех (ожидаемо)")
	fmt.Println("История email:", emailNotifier.Messages)
	fmt.Println("История push:", pushNotifier.Messages)
}

func showEmptyMessage(emailNotifier *practice.EmailNotifier, pushNotifier *practice.PushNotifier) {
	emailNotifierMessagesLenBeforeEmptyMessage := len(emailNotifier.Messages)

	allNotifiers := []practice.Notifier{emailNotifier, pushNotifier}
	res := practice.Dispatch(allNotifiers, "")

	if res == nil {
		fmt.Println("Случай с пустым message завершился успехом, внезапно")
		return
	}

	fmt.Println("Случай с пустым message завершился не успехом (ожидаемо)")

	if !errors.Is(res, practice.ErrEmptyMessage) {
		fmt.Printf("Не ожидаемая ошибка %v\n", res)
		return
	}

	fmt.Printf("Ожидаемая ошибка %v\n", res)

	if len(emailNotifier.Messages) == emailNotifierMessagesLenBeforeEmptyMessage {
		fmt.Println("История после ошибочного Dispatch сохранена (ожидаемо)")
		fmt.Println("Здесь можно проверять только emailNotifier тк мы знаем что по контракту Dispatch упадет на email и все, значит только эта история и могла пострадать")
	} else {
		fmt.Println("История после ошибочного Dispatch повреждена некорректным сообщением")
	}
}

func showMissingConfig() {
	emptyPushNotifier := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "Device"},
		DeviceID:     "",
	}

	allNotifiers := []practice.Notifier{emptyPushNotifier}
	res := practice.Dispatch(allNotifiers, "some message") // да я знаю что можно было написать Send

	if res == nil {
		fmt.Println("Случай с пустым deviceId завершился успехом, внезапно")
		return
	}

	fmt.Println("Случай с пустым deviceId завершился не успехом (ожидаемо)")

	var sendErr *practice.SendError
	if errors.As(res, &sendErr) {
		fmt.Printf(
			"ошибка %v - с полями Channel:%s, Reason:%s (ожидаемо)\n",
			res, sendErr.Channel, sendErr.Reason,
		)
	} else {
		fmt.Printf("Не ожидаемая ошибка %v\n", res)
	}
}

func showBrokenNotifier() {
	_, pushNotifier := newNotifiers()
	brokenNotifier := &practice.BrokenNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "Broken"},
	}

	allNotifiers := []practice.Notifier{brokenNotifier, pushNotifier}
	failures := practice.SafeDispatch(allNotifiers, "broken")

	if len(failures) == 0 {
		fmt.Println("Сломанный канал отработал без ошибки, внезапно")
	}

	for _, err := range failures {
		if errors.Is(err, practice.ErrNotifierPanic) {
			fmt.Printf("Канал сломался, но мы живы (ожидаемо): %v\n", err)
		} else {
			fmt.Printf("Не ожидаемая ошибка %v\n", err)
		}
	}

	if len(pushNotifier.Messages) == 1 {
		fmt.Println("Push после паники получил сообщение (ожидаемо)")
	} else {
		fmt.Println("Push после паники не получил сообщение, внезапно")
	}

	fmt.Println("История push:", pushNotifier.Messages)
}
