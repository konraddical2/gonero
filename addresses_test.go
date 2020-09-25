package gonero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddresses(t *testing.T) {
	type validTest struct {
		addr      string
		net       string
		spendKey  string
		viewKey   string
		paymentID string
		//privViewKey  string
		//privSpendKey string
	}

	validTests := []validTest{
		{"55hKAMnUWXaWXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmCKyGuXY", "stage",
			"65fd9e473d77c9b0894118f312b7f94931079108a12018720840696fcd1f9bb7",
			"c97cdd4e0b7985be9e436a6f360947c795cb4e890bd89ca7e620a5960efd8864", ""},
		{"481SgRxo8hwBCY4z6r88JrN5X8JFCJYuJUDuJXGybTwaVKyoJPKoGj3hQRAEGgQTdmV1xH1URdnHkJv6He5WkEbq6iKhr94", "main",
			"a87de9f3e64a223cf9a5abb4de0c0f7e044384698f0c0da2c09aa432912135a9",
			"5b80f6a85971f2f18ee232e161f724a77fd22afb6d911f6b1be5484eafa30032", ""},
		{"9wviCeWe2D8XS82k2ovp5EUYLzBt9pYNW2LXUFsZiv8S3Mt21FZ5qQaAroko1enzw3eGr9qC7X1D7Geoo2RrAotYPwq9Gm8", "test",
			"7d996b0f2db6dbb5f2a086211f2399a4a7479b2c911af307fdc3f7f61a88cb0e",
			"1c06bcac7082f73af10460b5f2849aded79374b2fbdaae5d9384b9b6514fddcb", ""},
	}

	t.Run("valid addresses", func(t *testing.T) {
		for _, tc := range validTests {
			addr := NewAddress(tc.addr)
			assert.True(t, addr.Valid())
			assert.Equal(t, tc.net, addr.Net())
			assert.Equal(t, tc.viewKey, addr.ViewKey())
			assert.Equal(t, tc.spendKey, addr.SpendKey())
			switch tc.net {
			case "main":
				assert.True(t, addr.IsMainnet())
			case "test":
				assert.True(t, addr.IsTestnet())
			case "stage":
				assert.True(t, addr.IsStagenet())
			}
			iaddr := addr.WithPaymentID("dead")
			assert.Equal(t, "dead", iaddr.PaymentID())
			assert.Equal(t, tc.addr, iaddr.BaseAddress().Addr)

			assert.False(t, addr.CheckPrivateSpendKey("zzz"), "CheckPrivateViewKey returns false when given private key with invalid hex")
			assert.False(t, addr.CheckPrivateViewKey("zzz"), "CheckPrivateSpendKey returns false when given private key with invalid hex")

			//if tc.privViewKey != "" {
			//	assert.True(t, addr.CheckPrivateViewKey(tc.privViewKey), "CheckPrivateViewKey returns true when given valid private key")
			//}
			//if tc.privSpendKey != "" {
			//	assert.True(t, addr.CheckPrivateSpendKey(tc.privSpendKey), "CheckPrivateSpendKey returns true when given valid private key")
			//}
		}
	})

	validTests = []validTest{
		{"731FGTtRm1j1iz4jSkGP7WY5MitM2vKGpb3pKGRk7NGYDcdZk96jgKQ1rNXTL2JT47FpE24VZok8UbBbtGfxpTBNHUvMdyL", "stage",
			"14085816c3ff4a0453e2813100c239b9c939af510fdd35cb901e9d702192d94b",
			"6a8410a434dc470516d27ad9900774589182d9ab7551adcc5d9213db308cc992", ""},
	}

	t.Run("valid subaddresses", func(t *testing.T) {
		for _, tc := range validTests {
			addr := NewSubAddress(tc.addr)
			assert.True(t, addr.Valid())
			assert.Equal(t, tc.net, addr.Net())
			assert.Equal(t, tc.viewKey, addr.ViewKey())
			assert.Equal(t, tc.spendKey, addr.SpendKey())
			switch tc.net {
			case "main":
				assert.True(t, addr.IsMainnet())
			case "test":
				assert.True(t, addr.IsTestnet())
			case "stage":
				assert.True(t, addr.IsStagenet())
			}
		}
	})

	validTests = []validTest{
		{"5F6WcXdbGMqLmTgH5RRZUn2khBh3VTV6vDqqY5QzdhrTRHWbKora7uGbiX9PHztEzEWv6LnShypfMFNVj7ZhzEFC7BSVDzdCRR5Fak3C1w", "stage",
			"", "", "f29f852b820d6c81"},
	}
	t.Run("valid integrated addresses", func(t *testing.T) {
		for _, tc := range validTests {
			addr := NewIntegratedAddress(tc.addr)
			assert.True(t, addr.Valid())
			assert.Equal(t, tc.net, addr.Net())
			assert.Equal(t, tc.paymentID, addr.PaymentID())
			switch tc.net {
			case "main":
				assert.True(t, addr.IsMainnet())
			case "test":
				assert.True(t, addr.IsTestnet())
			case "stage":
				assert.True(t, addr.IsStagenet())
			}
		}
	})

	type invalidTest struct {
		addr string
	}

	invalidTests := []invalidTest{
		{"foobar"},
		{"O(*&(*&(KJLDFNDSF&KJ"},
		{"9wviCeWe2D8XS82k2ovp5EUYLzBt9pYNW2LXUFsZiv8S3Mt21FZ5qQaAroko1enzw3eGr9qC7X1D7Geoo2RrAotYPwq9Gm7"},
	}
	t.Run("invalid addresses", func(t *testing.T) {
		for _, tc := range invalidTests {
			addr := NewAddress(tc.addr)
			assert.Nil(t, addr)
			//assert.False(t, addr.Valid())
			//assert.Equal(t, "", addr.Net())
			//assert.Equal(t, "", addr.ViewKey())
			//assert.Equal(t, "", addr.SpendKey())
		}
	})
}

func TestPaymentIds(t *testing.T) {
	type test struct {
		addr      string
		payID     string
		integAddr string
	}

	tests := []test{
		{"55hKAMnUWXaWXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmCKyGuXY", "dead",
			"5FPzBAby7o6WXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmHj8GPitKopZLZrtqSQ"},
		{"55hKAMnUWXaWXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmCKyGuXY", "beef",
			"5FPzBAby7o6WXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmHj8GPitKop1U2755t3"},
	}

	for _, tc := range tests {
		addr := NewAddress(tc.addr)
		integAddr := addr.WithPaymentID(tc.payID)
		assert.Equal(t, tc.integAddr, integAddr.Addr)
		assert.Equal(t, tc.payID, integAddr.PaymentID())
	}
}
