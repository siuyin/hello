# Playwrite is a web crawler
Similar to Selenium

## Missing dependencies
sudo apt install libglib2.0 libnss3 libatk1.0 libatk-bridge2.0 libcups2 libdrm-amdgpu1 libxkbcommon-x11-0 libxcomposite-dev libxdamage-dev libxrandr2  libgbm-dev libgtk-3-0 libasound2 libxshmfence1

## Important browser launch options
```
	opt := playwright.BrowserTypeLaunchOptions{
		Args: []string{"--disable-gpu"},
	}
	browser, err := pw.Chromium.Launch(opt)
```
