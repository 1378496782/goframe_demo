async function loadUserProfile() {
    const errorEl = document.getElementById('errorMessage');
    const profileInfo = document.getElementById('profileInfo');

    try {
        const response = await apiRequest('/user/profile', {
            method: 'GET'
        });

        if (response.code === 0 && response.data) {
            const user = response.data;
            errorEl.style.display = 'none';
            profileInfo.style.display = 'block';

            document.getElementById('userId').textContent = user.id;
            document.getElementById('userPassport').textContent = user.passport;
            document.getElementById('userNickname').textContent = user.nickname;
            document.getElementById('userCreateAt').textContent = user.createAt || '-';
            document.getElementById('userUpdateAt').textContent = user.updateAt || '-';
        } else {
            profileInfo.style.display = 'none';
            errorEl.textContent = response.message || '加载失败';
            errorEl.style.display = 'block';
        }
    } catch (error) {
        console.error('加载用户资料失败:', error);
        profileInfo.style.display = 'none';
        errorEl.textContent = '加载失败，请稍后再试';
        errorEl.style.display = 'block';
    }
}

document.addEventListener('DOMContentLoaded', loadUserProfile);
