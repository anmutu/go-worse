package ternary

// Ternary 实现三元表达式功能
func Ternary[T any](condition bool, trueVal, falseVal T) T {
    if condition {
        return trueVal
    }
    return falseVal
}