// API 基础路径
const API_BASE = '/api';

// 通用 API 请求函数
async function apiRequest(url, options = {}) {
    try {
        const response = await fetch(url, {
            ...options,
            credentials: 'include', // 携带 cookies
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            }
        });
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('API 请求错误:', error);
        throw error;
    }
}

// 检查登录状态
async function checkLoginStatus() {
    try {
        const response = await apiRequest('/user/is-signed-in', {
            method: 'POST'
        });
        if (response.code === 0) {
            return response.data?.OK === true;
        }
        return false;
    } catch (error) {
        console.error('检查登录状态失败:', error);
        return false;
    }
}

// 更新导航栏显示
async function updateNavForAuth() {
    const isLoggedIn = await checkLoginStatus();
    const loginLink = document.getElementById('loginLink');
    const signupLink = document.getElementById('signupLink');
    const profileLink = document.getElementById('profileLink');
    const logoutBtn = document.getElementById('logoutBtn');

    if (loginLink) loginLink.style.display = isLoggedIn ? 'none' : 'inline';
    if (signupLink) signupLink.style.display = isLoggedIn ? 'none' : 'inline';
    if (profileLink) profileLink.style.display = isLoggedIn ? 'inline' : 'none';
    if (logoutBtn) logoutBtn.style.display = isLoggedIn ? 'inline' : 'none';

    return isLoggedIn;
}

// 登出
async function logout() {
    try {
        const response = await apiRequest('/user/sign-out', {
            method: 'POST'
        });
        if (response.code === 0) {
            window.location.href = 'index.html';
        }
    } catch (error) {
        console.error('登出失败:', error);
    }
}

// 页面加载时执行
document.addEventListener('DOMContentLoaded', async () => {
    // 更新导航栏
    await updateNavForAuth();

    // 绑定登出按钮事件
    const logoutBtn = document.getElementById('logoutBtn');
    if (logoutBtn) {
        logoutBtn.addEventListener('click', logout);
    }
});

// 显示错误消息
function showError(elementId, message) {
    const errorEl = document.getElementById(elementId);
    const successEl = document.getElementById('successMessage');
    if (errorEl) {
        errorEl.textContent = message;
        errorEl.style.display = 'block';
    }
    if (successEl) {
        successEl.style.display = 'none';
    }
}

// 显示成功消息
function showSuccess(elementId, message) {
    const errorEl = document.getElementById('errorMessage');
    const successEl = document.getElementById(elementId);
    if (errorEl) {
        errorEl.style.display = 'none';
    }
    if (successEl) {
        successEl.textContent = message;
        successEl.style.display = 'block';
    }
}

// 隐藏所有消息
function hideMessages() {
    const errorEl = document.getElementById('errorMessage');
    const successEl = document.getElementById('successMessage');
    if (errorEl) {
        errorEl.style.display = 'none';
    }
    if (successEl) {
        successEl.style.display = 'none';
    }
}

