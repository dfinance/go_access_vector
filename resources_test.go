package glav

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestOraclePath(t *testing.T) {
	access := OracleAccessVector("DFI", "BTC")
	expected, _ := hex.DecodeString("01b1a724361d17c9866b12e199ecdb17eb5cb16630b647bbc997fe65362920e3bb")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}

	access = OracleAccessVector("BTC", "ETH")
	expected, _ = hex.DecodeString("01a7183ec0c4d32fd9a2705e1e6844035c5238598bf45167742e9db3735af96cc1")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}

func TestBlockMetadataVector(t *testing.T) {
	access := BlockMetadataVector()
	expected, _ := hex.DecodeString("01ada6f79e8eddfdf986687174de1000df3c5fa45e9965ece812fed33332ec543a")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}

func TestTimeMetadataVector(t *testing.T) {
	access := TimeMetadataVector()
	expected, _ := hex.DecodeString("01bb4980dfba817aaa64cb4b3a75ee1e1d8352111dead64c5c4f05fb7b4c85bb3e")
	if !reflect.DeepEqual(access, expected) {
		t.Errorf("Expected %X path. Actual: %X", expected, access)
	}
}

func TestBalanceVector(t *testing.T) {
	t.Run("eth denom", func(t *testing.T) {
		access := BalanceVector("eth")
		expected, _ := hex.DecodeString("0138f4f2895881c804de0e57ced1d44f02e976f9c6561c889f7b7eef8e660d2c9a")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})

	t.Run("dfi denom", func(t *testing.T) {
		access := BalanceVector("dfi")
		expected, _ := hex.DecodeString("01608540feb9c6bd277405cfdc0e9140c1431f236f7d97865575e830af3dd67e7e")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})
}

func TestCurrencyInfoVector(t *testing.T) {
	t.Run("eth denom", func(t *testing.T) {
		access := CurrencyInfoVector("eth")
		expected, _ := hex.DecodeString("012a00668b5325f832c28a24eb83dffa8295170c80345fbfbf99a5263f962c76f4")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})

	t.Run("dfi denom", func(t *testing.T) {
		access := CurrencyInfoVector("dfi")
		expected, _ := hex.DecodeString("01f3a1f15d7b13931f3bd5f957ad154b5cbaa0e1a2c3d4d967f286e8800eeb510d")
		if !reflect.DeepEqual(access, expected) {
			t.Errorf("Expected %X path. Actual: %X", expected, access)
		}
	})
}
