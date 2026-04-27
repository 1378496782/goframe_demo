document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('loginForm');
    const togglePassword = document.getElementById('togglePassword');
    const passwordInput = document.getElementById('password');
    const loginBtn = document.getElementById('loginBtn');

    // 密码显示/隐藏功能
    togglePassword.addEventListener('click', () => {
        const type = passwordInput.getAttribute('type') === 'password' ? 'text' : 'password';
        passwordInput.setAttribute('type', type);
        togglePassword.textContent = type === 'password' ? '👁️' : '🙈';
    });

    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        hideMessages();

        // 添加加载状态
        loginBtn.classList.add('loading');
        loginBtn.disabled = true;

        const formData = {
            Passport: document.getElementById('passport').value,
            Password: document.getElementById('password').value
        };

        try {
            const response = await apiRequest('/user/sign-in', {
                method: 'POST',
                body: JSON.stringify(formData)
            });

            if (response.code === 0) {
                showSuccess('successMessage', '登录成功！2秒后跳转到首页...');
                setTimeout(() => {
                    window.location.href = 'index.html';
                }, 2000);
            } else {
                showError('errorMessage', response.message || '登录失败');
            }
        } catch (error) {
            showError('errorMessage', '网络错误，请稍后再试');
        } finally {
            // 移除加载状态
            loginBtn.classList.remove('loading');
            loginBtn.disabled = false;
        }
    });

    // 添加键盘快捷方式
    document.addEventListener('keydown', (e) => {
        if (e.ctrlKey && e.key === 'Enter') {
            loginForm.dispatchEvent(new Event('submit'));
        }
    });
});
