// Package filemonitor
// Summary: Directory watcher with debounce and grouping.
// Details: 基于 fsnotify 监听数据目录，支持按分组筛选（正则/排除），并带抖动窗口与回调，
// 主要用于自动解密触发。
package filemonitor

