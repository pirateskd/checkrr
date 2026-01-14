func (c *Checkrr) isCorrupt(path string) (bool, string) {
	switch c.ScanMode {
	case "fast":
		// current behavior: ffprobe only
		return c.fastProbe(path)

	case "enhanced":
		// decode up to N seconds or until first decode error
		return c.enhancedDecode(path, time.Duration(c.EnhancedSeconds)*time.Second)

	case "full":
		// full decode
		return c.fullDecode(path)

	default:
		return c.fastProbe(path)
	}
}
