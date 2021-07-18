package pointersanderrors_test

import (
	"errors"
	"testing"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/pointersanderrors"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := pointersanderrors.Wallet{}
		wallet.Deposit(pointersanderrors.Bitcoin(10))

		assertBalance(t, wallet.Balance(), pointersanderrors.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := pointersanderrors.NewWallet(pointersanderrors.Bitcoin(20))

		err := wallet.Withdraw(pointersanderrors.Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet.Balance(), pointersanderrors.Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := pointersanderrors.NewWallet(pointersanderrors.Bitcoin(20))

		err := wallet.Withdraw(pointersanderrors.Bitcoin(100))

		assertError(t, err, errors.New("cannot withdraw, insufficient funds"))
		assertBalance(t, wallet.Balance(), pointersanderrors.Bitcoin(20))
	})

}

func assertBalance(t testing.TB, got, want pointersanderrors.Bitcoin) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q, got, want", got, want)
	}
}

func TestBitcoinStringer(t *testing.T) {
	got := pointersanderrors.Bitcoin(20).String()

	want := "20 BTC"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
