# Deployment Guide

## Quick Deploy to GitHub Pages

### 1. Initialize and commit (if not done already)
```bash
git init
git add .
git commit -m "Initial Jekyll blog with syntax highlighting"
```

### 2. Push to GitHub
```bash
# Replace with your actual GitHub username
git remote add origin https://github.com/LucasKonrath/devportal.git
git branch -M main
git push -u origin main
```

### 3. Enable GitHub Pages
1. Go to https://github.com/LucasKonrath/devportal
2. Click **Settings**
3. Click **Pages** in left sidebar
4. Under **Source**, select **main** branch
5. Click **Save**

### 4. Wait for build
- GitHub will automatically build your Jekyll site
- Check the Actions tab to see build progress
- Your site will be live at: https://LucasKonrath.github.io/devportal/

### 5. Test locally first (optional)
```bash
bundle install
bundle exec jekyll serve
# Visit http://localhost:4000/devportal/
```

## Configuration Already Done ✅
- `_config.yml` has been updated with:
  - `url: "https://LucasKonrath.github.io"`
  - `baseurl: "/devportal"`
- Syntax highlighting is configured
- All images are in the correct location

## Future Updates
To publish new posts:
```bash
# Create new post in _posts/YYYY-MM-DD-title.md
git add _posts/YYYY-MM-DD-your-new-post.md
git commit -m "Add new post"
git push
```

GitHub Pages will automatically rebuild your site!
