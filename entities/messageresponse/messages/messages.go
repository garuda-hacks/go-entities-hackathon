package messages

// Message wrapper.
type Message struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//Code 34000 - 3499 User & Auth bn;
var TelErrUserNotFound = Message{Code: "34000", Message: "Pengguna tidak ditemukan pada sistem"}
var TelErrUserRoleExists = Message{Code: "34001", Message: "User sudah di assign untuk role ini. Mohon periksa kembali data Anda"}
var TelErrNonActiveUser = Message{Code: "34002", Message: "Pengguna tidak terdaftar Pada Sistem"}
var TelErrPassword = Message{Code: "34003", Message: "Kata sandi yang di masukkan tidak sesuai"}
var TelErrUserSave = Message{Code: "34004", Message: "Pengguna tidak dapat di simpan. Mohon periksa kembali data Anda"}
var TelErrCodeNotValid = Message{Code: "34005", Message: "Activation Code Tidak Valid"}
var TelErrUserIsActive = Message{Code: "34006", Message: "User Sudah Aktif!"}
var TelErrEmailAlreadyUsed = Message{Code: "34007", Message: "Email sudah ada yang menggunakan"}
var TelErrUserNonActive = Message{Code: "34008", Message: "User belum aktif"}
var TelErrUserUnAuthorize = Message{Code: "34009", Message: "Unauthorized Access"}
var TelErrInvalidBind = Message{Code: "34010", Message: "Permintaan tidak dapat diproses"}
var TelErrHandphoneAlreadyUsed = Message{Code: "34011", Message: "Handphone sudah ada yang menggunakan"}
var TelErrDataCantDecrypt = Message{Code: "34012", Message: "Data user cannot decrypt"}
var TelErrShouldBindError = Message{Code: "34013", Message: "Function Should Bind Error"}
var TelErrTripleDes = Message{Code: "34014", Message: "Function TripleDes Error"}
var TelErrValidate = Message{Code: "34015", Message: "Function Validate Error"}
var TelErrVerifyEmail = Message{Code: "34016", Message: "Function Verify Email Error"}
var TelSuccessLogin = Message{Code: "34017", Message: "Login Success"}

// Code 39000 - 39999 Server error
var TelErrRevocerRoute = Message{Code: "39000", Message: "Terjadi kesalahan routing"}
var TelErrPageNotFound = Message{Code: "39404", Message: "Halaman Tidak ditemukan"}
