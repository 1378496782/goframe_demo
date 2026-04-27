async function loadUserList() {
    const tbody = document.getElementById('userTableBody');

    try {
        const response = await apiRequest('/user/list', {
            method: 'GET'
        });

        if (response.code === 0 && response.data?.userList) {
            const users = response.data.userList;
            
            if (users.length === 0) {
                tbody.innerHTML = `
                    <tr>
                        <td colspan="4" class="loading-cell">暂无用户数据</td>
                    </tr>
                `;
                return;
            }

            tbody.innerHTML = users.map(user => `
                <tr>
                    <td>${user.id}</td>
                    <td>${user.passport}</td>
                    <td>${user.nickname}</td>
                    <td>${user.createAt || '-'}</td>
                </tr>
            `).join('');
        } else {
            tbody.innerHTML = `
                <tr>
                    <td colspan="4" class="loading-cell">${response.message || '加载失败'}</td>
                </tr>
            `;
        }
    } catch (error) {
        console.error('加载用户列表失败:', error);
        tbody.innerHTML = `
            <tr>
                <td colspan="4" class="loading-cell">加载失败，请稍后再试</td>
            </tr>
        `;
    }
}

document.addEventListener('DOMContentLoaded', loadUserList);
