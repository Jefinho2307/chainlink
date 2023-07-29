package testutils

var (
	DefaultSecretsBytes      = []byte{0xaa, 0xbb, 0xcc}
	DefaultSecretsBase64     = "qrvM"
	DefaultSecretsUrlsBytes  = []byte{0x01, 0x02, 0x03}
	DefaultSecretsUrlsBase64 = "AQID"
	DefaultArg1              = "arg1"
	DefaultArg2              = "arg2"

	// This has been generated using the Functions Toolkit (https://github.com/smartcontractkit/functions-toolkit/blob/main/src/SecretsManager.ts) and decrypts to the JSON string `{"0x0":"qrvM"}`
	DefaultThresholdSecretsHex = "0x7b225444483243747874223a2265794a48636d393163434936496c41794e5459694c434a44496a6f69533035305a5559325448593056553168543341766148637955584a4b65545a68626b3177527939794f464e78576a59356158646d636a6c4f535430694c434a4d59574a6c62434936496b464251554642515546425155464251554642515546425155464251554642515546425155464251554642515546425155464251554642515545394969776956534936496b4a45536c6c7a51334e7a623055334d6e6444574846474e557056634770585a573157596e565265544d796431526d4d32786c636c705a647a4671536e6c47627a5256615735744e6d773355456855546e6b7962324e746155686f626c51354d564a6a4e6e5230656c70766147644255326372545430694c434a5658324a6863694936496b4a4961544e69627a5a45536d396a4d324d344d6c46614d5852724c325645536b4a484d336c5a556d783555306834576d684954697472623264575a306f33546e4e456232314b5931646853544979616d63305657644f556c526e57465272655570325458706952306c4a617a466e534851314f4430694c434a46496a6f694d7a524956466c354d544e474b307836596e5a584e7a6c314d6d356c655574514e6b397a656e467859335253513239705a315534534652704e4430694c434a47496a6f69557a5132596d6c6952545a584b314176546d744252445677575459796148426862316c6c6330684853556869556c56614e303155556c6f345554306966513d3d222c2253796d43747874223a2253764237652f4a556a552b433358757873384e5378316967454e517759755051623730306a4a6144222c224e6f6e6365223a224d31714b557a6b306b77374767593538227d"
)
