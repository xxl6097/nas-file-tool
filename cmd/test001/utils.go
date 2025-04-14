package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tips(title string) {
	str := strings.ReplaceAll(title, "请输入", "")
	str = strings.ReplaceAll(str, "please input", "")
	str = strings.ReplaceAll(str, "：", "")
	str = strings.ReplaceAll(str, ":", "")
	str = fmt.Sprintf("【%s】不允许输入空", str)
	fmt.Println(str)
}
func InputStringEmpty(title, defaultString string) string {
	reader := bufio.NewReader(os.Stdin)
	//glog.Print(title)
	fmt.Print(title)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return InputString(title)
	}
	if input == "" {
		return defaultString
	}
	//return strings.TrimSpace(input)
	return input
}

func InputString(title string) string {
	reader := bufio.NewReader(os.Stdin)
	//glog.Print(title)
	fmt.Print(title)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return InputString(title)
	}
	//return strings.TrimSpace(input)
	if len(input) == 0 {
		tips(title)
		return InputString(title)
	}
	return input
}
func InputInt(title string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(title)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return InputInt(title)
	}
	if len(input) == 0 {
		tips(title)
		return InputInt(title)
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return InputInt(title)
	}
	return num
}

	hash.Reset()
	return result, hex.EncodeToString(result)
}

func FormatSize(size int64) string {
	sizes := []string{`B`, `KB`, `MB`, `GB`, `TB`, `PB`, `EB`, `ZB`, `YB`}
	i := 0
	for size >= 1024 && i < len(sizes)-1 {
		size /= 1024
		i++
	}
	return fmt.Sprintf(`%d%s`, size, sizes[i])
}

func BytesToString(b []byte, r ...int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bytesPtr := sh.Data
	bytesLen := sh.Len
	switch len(r) {
	case 1:
		r[0] = If(r[0] > bytesLen, bytesLen, r[0])
		bytesLen -= r[0]
		bytesPtr += uintptr(r[0])
	case 2:
		r[0] = If(r[0] > bytesLen, bytesLen, r[0])
		bytesLen = If(r[1] > bytesLen, bytesLen, r[1]) - r[0]
		bytesPtr += uintptr(r[0])
	}
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: bytesPtr,
		Len:  bytesLen,
	}))
}

func StringToBytes(s string, r ...int) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	strPtr := sh.Data
	strLen := sh.Len
	switch len(r) {
	case 1:
		r[0] = If(r[0] > strLen, strLen, r[0])
		strLen -= r[0]
		strPtr += uintptr(r[0])
	case 2:
		r[0] = If(r[0] > strLen, strLen, r[0])
		strLen = If(r[1] > strLen, strLen, r[1]) - r[0]
		strPtr += uintptr(r[0])
	}
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: strPtr,
		Len:  strLen,
		Cap:  strLen,
	}))
}

func GetSlicePrefix[T any](data *[]T, n int) *[]T {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(data))
	return (*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sliceHeader.Data,
		Len:  n,
		Cap:  n,
	}))
}

func GetSliceSuffix[T any](data *[]T, n int) *[]T {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(data))
	return (*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sliceHeader.Data + uintptr(sliceHeader.Len-n),
		Len:  n,
		Cap:  n,
	}))
}

func GetSliceChunk[T any](data *[]T, start, end int) *[]T {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(data))
	return (*[]T)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sliceHeader.Data + uintptr(start),
		Len:  end - start,
		Cap:  end - start,
	}))
}

func CheckBinaryPack(data []byte) (byte, byte, bool) {
	if len(data) >= 8 {
		if bytes.Equal(data[:4], []byte{34, 22, 19, 17}) {
			if data[4] == 20 || data[4] == 21 {
				return data[4], data[5], true
			}
		}
	}
	return 0, 0, false
}

func BytesToHexString(bytes []byte) string {
	var hexValues []string
	for _, b := range bytes {
		// 将每个字节转换为十六进制字符串，并添加 0x 前缀
		hexValues = append(hexValues, fmt.Sprintf("0x%02x", b))
	}
	// 使用逗号连接所有十六进制字符串
	return strings.Join(hexValues, ", ")
}

// DivideAndCeil 函数用于进行除法并向上取整
func DivideAndCeil(a, b int) int {
	// 将整数转换为 float64 类型进行除法运算
	result := float64(a) / float64(b)
	// 使用 math.Ceil 函数进行向上取整
	result = math.Ceil(result)
	// 将结果转换回整数类型
	return int(result)
}

func Divide(a, b int) int {
	return DivideAndCeil(a, b) * b
}

// IsWindows 判断是否为 Windows 系统
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func GetDataByJson[T any](r *http.Request) (*T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func BindJSON[T any](r *http.Request) (*T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func IsMacOs() bool {
	if strings.Compare(runtime.GOOS, "darwin") == 0 {
		return true
	}
	return false
}

func IsLinux() bool {
	if strings.Compare(runtime.GOOS, "linux") == 0 {
		return true
	}
	return false
}

func SplitVersion(v string) []string {
	// 去除前缀标识（如 "v1.2.3" → "1.2.3"）
	v = strings.TrimLeft(v, "v")
	return strings.Split(v, ".")
}

// CompareVersions 0:相等；1：v1>v2;-1:v1<v2
func CompareVersions(v1, v2 string) int {
	seg1 := SplitVersion(v1)
	seg2 := SplitVersion(v2)
	maxLen := int(math.Max(float64(len(seg1)), float64(len(seg2))))

	for i := 0; i < maxLen; i++ {
		num1 := getSegmentValue(seg1, i)
		num2 := getSegmentValue(seg2, i)

		if num1 > num2 {
			return 1 // v1 > v2
		} else if num1 < num2 {
			return -1 // v1 < v2
		}
	}
	return 0 // 相等
}

func getSegmentValue(seg []string, idx int) int {
	if idx >= len(seg) {
		return 0 // 自动补零处理长度不一致情况
	}
	num, _ := strconv.Atoi(seg[idx])
	return num
}

func ReplaceNewVersionBinName(filename, v string) string {
	re := regexp.MustCompile(`_v\d+\.\d+\.\d+_`)
	newName := re.ReplaceAllString(filename, fmt.Sprintf("_%s_", v)) // 替换为单个下划线
	fmt.Println(newName)
	return newName
}

func GetSelfSize() uint64 {
	// 获取当前可执行文件的路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("获取可执行文件路径时出错: %v\n", err)
		return 0
	}
	// 获取文件信息
	fileInfo, err := os.Stat(exePath)
	if err != nil {
		fmt.Printf("获取文件信息时出错: %v\n", err)
		return 0
	}

	// 获取文件大小
	fileSize := fileInfo.Size()
	fmt.Printf("本程序自身大小为: %d 字节\n", fileSize)
	return uint64(fileSize)
}
