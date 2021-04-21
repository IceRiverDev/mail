package email

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {
	t.Run("test send mail", func(t *testing.T) {
		fmt.Println("send mail passed...")
	})

	t.Run("test mail configs", func(t *testing.T) {
		fmt.Println("mail configs passed...")
	})
}

func TestName(t *testing.T) {
	t.Run("test name", func(t *testing.T) {
		fmt.Println("Hello World")
	})

	t.Run("test age", func(t *testing.T) {
		fmt.Println("Hello God")
	})
}
