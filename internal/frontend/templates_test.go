package frontend

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTemplates(t *testing.T) {
	tpl, err := NewTemplates()
	require.NoError(t, err)

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, "header.html", nil)
	require.NoError(t, err)

	assert.Equal(t, `
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
<link rel="stylesheet" type="text/css" href="data:text/css;%20charset=utf-8;base64,LyoqKioqKiogR2xvYmFsICoqKioqKiovCgpib2R5LApkaXYsCmRsLApkdCwKZGQsCnVsLApvbCwKbGksCmgxLApoMiwKaDMsCmg0LApoNSwKaDYsCnByZSwKY29kZSwKZm9ybSwKZmllbGRzZXQsCmxlZ2VuZCwKaW5wdXQsCmJ1dHRvbiwKdGV4dGFyZWEsCnAsCmJsb2NrcXVvdGUsCnRoLAp0ZCB7CiAgbWFyZ2luOiAwOwogIHBhZGRpbmc6IDA7Cn0KCmZpZWxkc2V0LAppbWcgewogIGJvcmRlcjogMDsKfQphZGRyZXNzLApjYXB0aW9uLApjaXRlLApjb2RlLApkZm4sCmVtLApzdHJvbmcsCnRoLAp2YXIsCm9wdGdyb3VwIHsKICBmb250LXN0eWxlOiBpbmhlcml0OwogIGZvbnQtd2VpZ2h0OiBpbmhlcml0Owp9CmRlbCwKaW5zIHsKICB0ZXh0LWRlY29yYXRpb246IG5vbmU7Cn0KbGkgewogIGxpc3Qtc3R5bGU6IG5vbmU7Cn0KY2FwdGlvbiwKdGggewogIHRleHQtYWxpZ246IGxlZnQ7Cn0KaDEsCmgyLApoMywKaDQsCmg1LApoNiB7CiAgZm9udC1zaXplOiAxMDAlOwogIGZvbnQtd2VpZ2h0OiBub3JtYWw7Cn0KcTpiZWZvcmUsCnE6YWZ0ZXIgewogIGNvbnRlbnQ6ICIiOwp9CmFiYnIsCmFjcm9ueW0gewogIGJvcmRlcjogMDsKICBmb250LXZhcmlhbnQ6IG5vcm1hbDsKfQpzdXAgewogIHZlcnRpY2FsLWFsaWduOiBiYXNlbGluZTsKfQpzdWIgewogIHZlcnRpY2FsLWFsaWduOiBiYXNlbGluZTsKfQpsZWdlbmQgewogIGNvbG9yOiAjMDAwOwp9CmlucHV0LApidXR0b24sCnRleHRhcmVhLApzZWxlY3QsCm9wdGdyb3VwLApvcHRpb24gewogIGZvbnQtZmFtaWx5OiBpbmhlcml0OwogIGZvbnQtc2l6ZTogaW5oZXJpdDsKICBmb250LXN0eWxlOiBpbmhlcml0OwogIGZvbnQtd2VpZ2h0OiBpbmhlcml0Owp9CmlucHV0LApidXR0b24sCnRleHRhcmVhLApzZWxlY3QgewogICpmb250LXNpemU6IDEwMCU7Cn0KCmFzaWRlLApkaWFsb2csCmZpZ3VyZSwKZm9vdGVyLApoZWFkZXIsCmhncm91cCwKbWVudSwKbmF2LApzZWN0aW9uIHsKICBkaXNwbGF5OiBibG9jazsKfQoKaHRtbCwKYm9keSB7CiAgbWFyZ2luOiAwOwogIHBhZGRpbmc6IDA7Cn0KCmh0bWwgewogIGJhY2tncm91bmQ6ICNmNmY5ZmM7Cn0KCmJvZHkgewogIGZvbnQtZmFtaWx5OiAtYXBwbGUtc3lzdGVtLCBCbGlua01hY1N5c3RlbUZvbnQsICJTZWdvZSBVSSIsIFJvYm90bywKICAgICJIZWx2ZXRpY2EgTmV1ZSIsIHNhbnMtc2VyaWY7CiAgLXdlYmtpdC1mb250LXNtb290aGluZzogYW50aWFsaWFzZWQ7CiAgLW1vei1vc3gtZm9udC1zbW9vdGhpbmc6IGdyYXlzY2FsZTsKfQoKLmlubmVyIHsKICB3aWR0aDogNjE0cHg7CiAgbWFyZ2luOiBhdXRvOwogIG1hcmdpbi1ib3R0b206IDJlbTsKfQoKLmJveCB7CiAgb3ZlcmZsb3c6IGhpZGRlbjsKICBib3JkZXItcmFkaXVzOiA0cHg7CiAgYm94LXNoYWRvdzogMCAxNXB4IDM1cHggcmdiYSg1MCwgNTAsIDkzLCAwLjEpLCAwIDVweCAxNXB4IHJnYmEoMCwgMCwgMCwgMC4wNyk7Cn0KCi5ib3gtaW5uZXIgewogIHBhZGRpbmc6IDM1cHg7Cn0KCi53aGl0ZSB7CiAgYmFja2dyb3VuZDogd2hpdGU7Cn0KCmgyIHsKICBmb250LXNpemU6IDEuNWVtOwogIGZvbnQtc3R5bGU6IG5vcm1hbDsKICBjb2xvcjogIzQ0NDsKICBtYXJnaW46IDAgMCAwLjhlbSAwOwogIHBhZGRpbmctYm90dG9tOiAwLjJlbTsKICBib3JkZXItYm90dG9tOiAxcHggc29saWQgI2VlZTsKfQoKdWwucGxhaW4gewogIGxpc3Qtc3R5bGU6IG5vbmU7CiAgLXdlYmtpdC1wYWRkaW5nLXN0YXJ0OiAwOwogIHBhZGRpbmctbGVmdDogMDsKfQoKYSB7CiAgY29sb3I6ICM2ZTQzZTg7CiAgdGV4dC1kZWNvcmF0aW9uOiBub25lOwp9CmE6aG92ZXIgewogIGNvbG9yOiAjMzIzMjVkOwp9CgovKioqKioqKiBTdGF0dXMvR3JhcGggQ29sb3JzICoqKioqKiovCgouc3RhdHVzLXVwIC5zdGF0dXMtdGltZSB7CiAgY29sb3I6ICMzZWNmOGU7Cn0KCi5zdGF0dXMtdXAgewogIGJhY2tncm91bmQ6ICMzZWNmOGU7Cn0KCi5zdGF0dXMtZG93biAuc3RhdHVzLXRpbWUgewogIGNvbG9yOiAjZmZlN2NiOwp9Cgouc3RhdHVzLWRvd24gewogIGJhY2tncm91bmQ6ICNlMjU5NTA7Cn0KCi8qKioqKioqIENsZWFyZml4ICoqKioqKiovCgouY2xlYXJmaXg6YWZ0ZXIgewogIHZpc2liaWxpdHk6IGhpZGRlbjsKICBkaXNwbGF5OiBibG9jazsKICBmb250LXNpemU6IDA7CiAgY29udGVudDogIiAiOwogIGNsZWFyOiBib3RoOwogIGhlaWdodDogMDsKfQoKKiBodG1sIC5jbGVhcmZpeCB7CiAgem9vbTogMTsKfSAvKiBJRTYgKi8KKjpmaXJzdC1jaGlsZCArIGh0bWwgLmNsZWFyZml4IHsKICB6b29tOiAxOwp9IC8qIElFNyAqLwoKLyoqKioqKiogSGVhZGVyICoqKioqKiovCgouaGVhZGVyIHsKICBwYWRkaW5nOiAzMHB4IDA7CiAgaGVpZ2h0OiA0MHB4OwogIHBvc2l0aW9uOiByZWxhdGl2ZTsKfQoKLmhlYWRlciBzcGFuIHsKICBjb2xvcjogIzZlNDNlODsKICBmb250LXNpemU6IDE2cHg7CiAgbGluZS1oZWlnaHQ6IDMxcHg7CiAgcG9zaXRpb246IGFic29sdXRlOwogIGxlZnQ6IDU0MHB4Owp9CgouaGVhZGluZyB7CiAgZmxvYXQ6IGxlZnQ7CiAgbWFyZ2luOiA3cHggMXB4Owp9Cgouc3RhdHVzZXMgewogIGZsb2F0OiBsZWZ0Owp9CgoubG9nbyB7CiAgZGlzcGxheTogaW5saW5lLWJsb2NrOwogIHBvc2l0aW9uOiByZWxhdGl2ZTsKICBiYWNrZ3JvdW5kOiB1cmwoLy5wb21lcml1bS9hc3NldHMvaW1nL2xvZ28tbG9uZy5zdmcpIG5vLXJlcGVhdDsKICB3aWR0aDogNjYzcHg7CiAgaGVpZ2h0OiAyNnB4OwogIGN1cnNvcjogcG9pbnRlcjsKfQoKLmxvZ286aG92ZXIgewogIG9wYWNpdHk6IDAuNzsKfQoKLyoqKioqKiogQ29udGVudCAqKioqKioqLwoubGFyZ2VzdGF0dXMgewogIGJvcmRlcjogMDsKICBwb3NpdGlvbjogcmVsYXRpdmU7CiAgei1pbmRleDogMTA7CgogIHBhZGRpbmc6IDAgMzZweDsKICBwYWRkaW5nLWxlZnQ6IDg0cHg7CgogIG1pbi1oZWlnaHQ6IDE1NXB4Owp9CgoubGFyZ2VzdGF0dXMgLnRpdGxlIHsKICBkaXNwbGF5OiBibG9jazsKICBwYWRkaW5nLXRvcDogNDZweDsKICBtYXJnaW4tYm90dG9tOiAxMHB4OwoKICBmb250LXNpemU6IDI5cHg7CiAgY29sb3I6ICMzMjMyNWQ7Cn0KCi5sYXJnZXN0YXR1cyAuc3RhdHVzLXRpbWUgewogIGRpc3BsYXk6IGJsb2NrOwogIGZvbnQtc2l6ZTogMTRweDsKICBjb2xvcjogIzg4OThhYTsKICBwYWRkaW5nLWJvdHRvbTogNDZweDsKfQoKLmNhdGVnb3J5IHsKICBtYXJnaW4tdG9wOiA0MHB4Owp9CgovKioqKioqKiBTdGF0dXNlcyAqKioqKioqLwoKLnN0YXR1c2VzIHsKICBmb250LXNpemU6IDAuN2VtOwp9Cgouc3RhdHVzLWJ1YmJsZSB7CiAgd2lkdGg6IDQ0cHg7CiAgaGVpZ2h0OiA0NHB4OwogIHBvc2l0aW9uOiBhYnNvbHV0ZTsKICBsZWZ0OiAyNHB4OwogIHRvcDogNTJweDsKICBiYWNrZ3JvdW5kLXBvc2l0aW9uOiBjZW50ZXI7CiAgYmFja2dyb3VuZC1yZXBlYXQ6IG5vLXJlcGVhdDsKICBib3JkZXItcmFkaXVzOiA1MCU7Cn0KCi50aXRsZS13cmFwcGVyIHsKICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7CiAgY29sb3I6ICMzMzM7CiAgbWluLWhlaWdodDogMTU1cHg7Cn0KCi5zdGF0dXMtdGltZSB7CiAgY29sb3I6ICM5OTk7Cn0KCi8qKioqKioqIGNhdGVnb3J5ICoqKioqKiovCi5jYXRlZ29yeSB7CiAgYmFja2dyb3VuZDogd2hpdGU7Cn0KCmRpdi5jYXRlZ29yeS1oZWFkZXIgewogIG1hcmdpbi1ib3R0b206IDI1cHg7Cn0KCi5jYXRlZ29yeS10aXRsZSB7CiAgY29sb3I6ICM1MjVmN2Y7CiAgZm9udC1zaXplOiAxNXB4OwogIHBhZGRpbmctcmlnaHQ6IDEwcHg7Cn0KCi5jYXRlZ29yeS1pY29uIHsKICBkaXNwbGF5OiBibG9jazsKICBwb3NpdGlvbjogcmVsYXRpdmU7CiAgdG9wOiAycHg7CiAgZmxvYXQ6IHJpZ2h0OwogIHdpZHRoOiAyN3B4OwogIGhlaWdodDogMjVweDsKICBiYWNrZ3JvdW5kOiB1cmwoLy5wb21lcml1bS9hc3NldHMvaW1nL2p3dC5zdmcpIDEwMCUgMCBuby1yZXBlYXQ7Cn0KCmRpdi5jYXRlZ29yeS1saW5rIHsKICBiYWNrZ3JvdW5kOiAjZjZmOWZjOwoKICBwYWRkaW5nOiAyNXB4IDM2cHg7CgogIGJvcmRlci1ib3R0b20tbGVmdC1yYWRpdXM6IDRweDsKICBib3JkZXItYm90dG9tLXJpZ2h0LXJhZGl1czogNHB4OwoKICBmb250LXNpemU6IDEzcHg7CgogIGNvbG9yOiAjNmI3YzkzOwp9CgovKiBGb290ZXIgKi8KZGl2I2Zvb3RlciB7CiAgLypiYWNrZ3JvdW5kOiByZ2JhKDAsMCwwLDAuMDUpOyovCiAgbWFyZ2luOiAwIDAgMDsKICBwYWRkaW5nOiA0MHB4IDAgMHB4OwogIHBvc2l0aW9uOiByZWxhdGl2ZTsKICBmb250LXNpemU6IDEzcHg7Cn0KCmRpdiNmb290ZXIgdWwgewogIHBhZGRpbmctbGVmdDogMTBweDsKfQpkaXYjZm9vdGVyIGxpIHsKICBkaXNwbGF5OiBpbmxpbmU7CiAgcGFkZGluZy1yaWdodDogMTVweDsKfQpkaXYjZm9vdGVyIGEgewogIGNvbG9yOiAjODg5OGFhOwogIHRleHQtZGVjb3JhdGlvbjogbm9uZTsKfQpkaXYjZm9vdGVyIGE6aG92ZXIgewogIGNvbG9yOiAjMzIzMjVkOwp9CgpkaXYjZm9vdGVyIGxpLmxhc3QgewogIGJvcmRlcjogMDsKfQoKZGl2I2Zvb3RlciBwIHsKICBjb2xvcjogIzg4OThhYTsKICBwb3NpdGlvbjogYWJzb2x1dGU7CiAgcmlnaHQ6IDEwcHg7CiAgdG9wOiA0MHB4Owp9CgovKiAtIFRhYmxlcyAqLwp0YWJsZSB0Ym9keSB0ciB0ZCB7CiAgYm9yZGVyLWNvbG9yOiAjNTI1ZjdmOwogIGJvcmRlci1zdHlsZTogc29saWQ7CiAgYm9yZGVyOiAwOwogIHBhZGRpbmc6IDE2cHggMTZweDsKfQp0YWJsZSB7CiAgd2lkdGg6IDEwMCU7CiAgYm9yZGVyLWNvbGxhcHNlOiBzZXBhcmF0ZTsKICBib3JkZXItc3BhY2luZzogMDsKICBtYXJnaW46IGF1dG87Cn0KdGFibGU6bm90KDpmaXJzdC1jaGlsZCkgewogIG1hcmdpbi10b3A6IDIwcHg7Cn0KdGFibGUgdGFibGUgewogIG1hcmdpbjogMTZweCAwIDhweCAwOwp9CnRhYmxlIHRoZWFkIHRyIHRoIHsKICBmb250LXdlaWdodDogNTAwOwogIGZvbnQtc2l6ZTogMTNweDsKICBjb2xvcjogIzZlNDNlODsKICB0ZXh0LXRyYW5zZm9ybTogdXBwZXJjYXNlOwogIHRleHQtYWxpZ246IGxlZnQ7CiAgcGFkZGluZzogMCAwIDhweCAxNnB4Owp9CnRhYmxlIHRoZWFkIHRyIHRoIHAgewogIGZvbnQtc2l6ZTogMTNweDsKCiAgdGV4dC10cmFuc2Zvcm06IHVwcGVyY2FzZTsKICBwYWRkaW5nOiAwOwogIGxpbmUtaGVpZ2h0OiAxNXB4Owp9CnRhYmxlIHRib2R5LAp0YWJsZSB0Ym9keSB0ZCA&#43;ICogewogIGZvbnQtc2l6ZTogMTRweDsKICBsaW5lLWhlaWdodDogMjBweDsKICB2ZXJ0aWNhbC1hbGlnbjogdG9wOwogIHBhZGRpbmctdG9wOiAwOwogIG92ZXJmbG93LXdyYXA6IGFueXdoZXJlOwp9Cgp0YWJsZSB0Ym9keSB0ciB0ZCB7CiAgYm9yZGVyLWNvbG9yOiByZ2IoMjI3LCAyMzIsIDIzOCk7CiAgYm9yZGVyLXN0eWxlOiBzb2xpZDsKICBwYWRkaW5nOiAxNnB4IDE2cHg7CiAgbWluLXdpZHRoOiA4MHB4Owp9CnRhYmxlIHRib2R5IHRyIHRkOmZpcnN0LWNoaWxkIHsKICBib3JkZXItbGVmdC13aWR0aDogMXB4Owp9CnRhYmxlIHRib2R5IHRyIHRkOmxhc3QtY2hpbGQgewogIGJvcmRlci1yaWdodC13aWR0aDogMXB4Owp9CnRhYmxlIHRib2R5IHRyOmZpcnN0LWNoaWxkID4gdGQgewogIGJvcmRlci10b3Atd2lkdGg6IDFweDsKfQp0YWJsZSB0Ym9keSB0ciB0ZCB7CiAgYm9yZGVyLWJvdHRvbS13aWR0aDogMXB4Owp9CnRhYmxlIHRib2R5IHRyOmZpcnN0LWNoaWxkID4gdGQ6Zmlyc3QtY2hpbGQgewogIGJvcmRlci10b3AtbGVmdC1yYWRpdXM6IDFweDsKfQp0YWJsZSB0Ym9keSB0cjpmaXJzdC1jaGlsZCA&#43;IHRkOmxhc3QtY2hpbGQgewogIGJvcmRlci10b3AtcmlnaHQtcmFkaXVzOiAxcHg7Cn0KdGFibGUgdGJvZHkgdHI6bGFzdC1jaGlsZCA&#43;IHRkOmZpcnN0LWNoaWxkIHsKICBib3JkZXItYm90dG9tLWxlZnQtcmFkaXVzOiAxcHg7Cn0KdGFibGUgdGJvZHkgdHI6bGFzdC1jaGlsZCA&#43;IHRkOmxhc3QtY2hpbGQgewogIGJvcmRlci1ib3R0b20tcmlnaHQtcmFkaXVzOiAxcHg7Cn0KdGFibGUgdGJvZHkgdHI6bnRoLWNoaWxkKDJuICsgMSkgdGQgewogIGJhY2tncm91bmQ6ICNmNmY5ZmM7Cn0KCmlucHV0LApidXR0b24sCmEuYnV0dG9uIHsKICBiYWNrZ3JvdW5kOiAjNmU0M2U4OwogIGJveC1zaGFkb3c6IDAgMnB4IDVweCAwIHJnYmEoNTAsIDUwLCA5MywgMC4xKSwgMCAxcHggMXB4IDAgcmdiYSgwLCAwLCAwLCAwLjA3KTsKICBib3JkZXItcmFkaXVzOiA0cHg7CiAgaGVpZ2h0OiAzMnB4OwogIGZvbnQtc2l6ZTogMTZweDsKICBjb2xvcjogI2Y2ZjlmYzsKICBmb250LXdlaWdodDogNTAwOwogIHBhZGRpbmc6IDAgMTJweDsKICBjdXJzb3I6IHBvaW50ZXI7CiAgb3V0bGluZTogbm9uZTsKICBkaXNwbGF5OiBpbmxpbmUtYmxvY2s7CiAgdGV4dC1kZWNvcmF0aW9uOiBub25lOwogIHRleHQtdHJhbnNmb3JtOiBub25lOwogIHRyYW5zaXRpb246IGJveC1zaGFkb3cgMTUwbXMgZWFzZS1pbi1vdXQ7Cn0K"/>
<link rel="icon" type="image/png" href="data:image/svg&#43;xml;base64,PHN2ZyBpZD0iTGF5ZXJfMSIgZGF0YS1uYW1lPSJMYXllciAxIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNjEuMTUgMTI0LjczIj48dGl0bGU&#43;bG9nby1vbmx5PC90aXRsZT48cGF0aCBkPSJNNzQyLjc2LDE2Mi44MmEyNC4zNiwyNC4zNiwwLDAsMC0yNC4zNi0yNC4zNkg1MDZhMjQuMzYsMjQuMzYsMCwwLDAtMjQuMzYsMjQuMzZWMjYzLjE5aDE2Ljgzdi0yOGgwYTM0LjExLDM0LjExLDAsMSwxLDY4LjIxLDBoMHYyOGgxMi4xOHYtMjhoMGEzNC4xMSwzNC4xMSwwLDEsMSw2OC4yMSwwaDB2MjhoMTIuMTh2LTI4aDBhMzQuMSwzNC4xLDAsMSwxLDY4LjIsMGgwdjI4aDE1LjMzWk00OTguNDQsMTg4LjEzYTM0LjExLDM0LjExLDAsMSwxLDY4LjIxLDBabTgwLjM5LDBhMzQuMTEsMzQuMTEsMCwxLDEsNjguMjEsMFptODAuMzksMGEzNC4xMSwzNC4xMSwwLDEsMSw2OC4yMSwwWiIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTQ4MS42MSAtMTM4LjQ2KSIvPjwvc3ZnPgo=" />
`, buf.String())
}
