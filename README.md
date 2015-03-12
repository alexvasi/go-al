go-al
=====

This is a wrapper around the most recent version of OpenAL-Soft. Right now it's only configured to compile on Windows and Linux, but patches for other platforms are welcome.

The primary difference between this library and Go OpenAL wrappers is that it doesn't wrap alh, which is deprecated. Instead, it provides examples showing you how to create your own context. In addition, it provides a bare bones .wav loader written in pure Go.
