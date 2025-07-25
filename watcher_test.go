package godev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContain(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		setup    func() *WatchHandler
		expected bool
	}{
		{
			name: "hidden file",
			path: ".gitignore",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{}
						},
					},
				}
			},
			expected: true,
		},
		{
			name: "unobserved file",
			path: "test/.git",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{".git"}
						},
					},
				}
			},
			expected: true,
		},
		{
			name: "observed file",
			path: "test/main.go",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{".git"}
						},
					},
				}
			},
			expected: false,
		},
		{
			name: "git folder in middle of path",
			path: "C:\\Users\\Cesar\\Packages\\Internal\\godev\\test\\manual\\.git\\objects\\pack",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{".git"}
						},
					},
				}
			},
			expected: true,
		},
		{
			name: "git folder in middle of path with unix style",
			path: "/Users/Cesar/Packages/Internal/godev/test/manual/.git/objects/pack",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{".git"}
						},
					},
				}
			},
			expected: true,
		},
		{
			name: "git folder in middle of path with project root",
			path: "test/manual/.git/objects/pack",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{".git"}
						},
					},
				}
			},
			expected: true,
		},
		{
			name: "git string in directory name but not excluded",
			path: "test/github-integration/code.go",
			setup: func() *WatchHandler {
				return &WatchHandler{
					WatchConfig: &WatchConfig{
						UnobservedFiles: func() []string {
							return []string{".git"}
						},
					},
				}
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := tt.setup()
			assert.Equal(t, tt.expected, handler.Contain(tt.path))
		})
	}
}
