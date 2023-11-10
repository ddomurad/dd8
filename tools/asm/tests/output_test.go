package tests

import (
	"strings"
	"testing"

	pkg "github.com/ddomurad/dd8/tools/asm/pkg"
	"github.com/ddomurad/dd8/tools/asm/pkg/output"
	"github.com/stretchr/testify/require"
)

func TestIntelHexFormat(t *testing.T) {
	t.Run("simple_test", func(t *testing.T) {
		byteCode := pkg.ByteCode{}
		err := byteCode.SetBytes(0x0000, []byte{0x01, 0x02, 0x03, 0x04})
		require.NoError(t, err)

		srec := output.GetIntelHEX(byteCode)
		srecLines := strings.Split(string(srec), "\n")

		require.Equal(t, ":0400000001020304F2", srecLines[0])
		require.Equal(t, ":00000001FF", srecLines[1])
	})

	t.Run("simple_test", func(t *testing.T) {
		byteCode := pkg.ByteCode{}
		err := byteCode.SetBytes(0x0000, []byte{0x01, 0x02, 0x03, 0x04})
		require.NoError(t, err)

		srec := output.GetIntelHEX(byteCode)
		srecLines := strings.Split(string(srec), "\n")

		require.Equal(t, ":0400000001020304F2", srecLines[0])
		require.Equal(t, ":00000001FF", srecLines[1])
	})

	t.Run("address_gap_tesp", func(t *testing.T) {
		byteCode := pkg.ByteCode{}
		err := byteCode.SetBytes(0x0000, []byte{0x01, 0x02, 0x03, 0x04})
		require.NoError(t, err)
		err = byteCode.SetBytes(0x1000, []byte{0x01, 0x02, 0x03, 0x04})
		require.NoError(t, err)

		srec := output.GetIntelHEX(byteCode)
		srecLines := strings.Split(string(srec), "\n")

		require.Equal(t, ":0400000001020304F2", srecLines[0])
		require.Equal(t, ":0410000001020304E2", srecLines[1])
		require.Equal(t, ":00000001FF", srecLines[2])
	})
}
