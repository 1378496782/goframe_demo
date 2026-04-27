document.addEventListener('DOMContentLoaded', () => {
    const signupForm = document.getElementById('signupForm');
    const togglePassword = document.getElementById('togglePassword');
    const togglePassword2 = document.getElementById('togglePassword2');
    const passwordInput = document.getElementById('password');
    const password2Input = document.getElementById('password2');
    const passwordStrengthBar = document.getElementById('passwordStrengthBar');
    const passwordStrengthText = document.getElementById('passwordStrengthText');
    const signupBtn = document.getElementById('signupBtn');

    // 密码显示/隐藏功能
    togglePassword.addEventListener('click', () => {
        const type = passwordInput.getAttribute('type') === 'password' ? 'text' : 'password';
        passwordInput.setAttribute('type', type);
        togglePassword.textContent = type === 'password' ? '👁️' : '🙈';
    });

    togglePassword2.addEventListener('click', () => {
        const type = password2Input.getAttribute('type') === 'password' ? 'text' : 'password';
        password2Input.setAttribute('type', type);
        togglePassword2.textContent = type === 'password' ? '👁️' : '🙈';
    });

    // 密码强度检测
    function checkPasswordStrength(password) {
        let strength = 0;
        if (password.length >= 6) strength++;
        if (password.length >= 10) strength++;
        if (/[A-Z]/.test(password)) strength++;
        if (/[0-9]/.test(password)) strength++;
        if (/[^A-Za-z0-9]/.test(password)) strength++;
        return strength;
    }

    passwordInput.addEventListener('input', () => {
        const password = passwordInput.value;
        const strength = checkPasswordStrength(password);
        const strengths = [
            { width: '0%', color: '#e0e0e0', text: '' },
            { width: '20%', color: '#ff6b6b', text: '密码强度：弱' },
            { width: '40%', color: '#ffa94d', text: '密码强度：一般' },
            { width: '60%', color: '#ffd43b', text: '密码强度：中等' },
            { width: '80%', color: '#69db7c', text: '密码强度：强' },
            { width: '100%', color: '#51cf66', text: '密码强度：非常强' }
        ];
        
        const current = strengths[strength];
        passwordStrengthBar.style.width = current.width;
        passwordStrengthBar.style.backgroundColor = current.color;
        passwordStrengthText.textContent = current.text;
        passwordStrengthText.style.color = current.color;
    });

    signupForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        hideMessages();

        // 验证密码一致性
        if (passwordInput.value !== password2Input.value) {
            showError('errorMessage', '两次输入的密码不一致');
            return;
        }

        // 添加加载状态
        signupBtn.classList.add('loading');
        signupBtn.disabled = true;

        const formData = {
            Passport: document.getElementById('passport').value,
            Password: document.getElementById('password').value,
            Password2: document.getElementById('password2').value,
            Nickname: document.getElementById('nickname').value
        };

        try {
            const response = await apiRequest('/user/sign-up', {
                method: 'POST',
                body: JSON.stringify(formData)
            });

            if (response.code === 0) {
                showSuccess('successMessage', '注册成功！2秒后跳转到登录页面...');
                setTimeout(() => {
                    window.location.href = 'login.html';
                }, 2000);
            } else {
                showError('errorMessage', response.message || '注册失败');
            }
        } catch (error) {
            showError('errorMessage', '网络错误，请稍后再试');
        } finally {
            // 移除加载状态
            signupBtn.classList.remove('loading');
            signupBtn.disabled = false;
        }
    });

    // 添加键盘快捷方式
    document.addEventListener('keydown', (e) => {
        if (e.ctrlKey && e.key === 'Enter') {
            signupForm.dispatchEvent(new Event('submit'));
        }
    });
});
