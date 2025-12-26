let inventory = [];
const elements = {
    form: document.getElementById('itemForm'),
    nameInput: document.getElementById('nameField'),
    qtyInput: document.getElementById('qtyField'),
    grid: document.getElementById('inventoryGrid'),
    editOverlay: document.getElementById('editPanel'),
    modifyForm: document.getElementById('modifyForm'),
    itemId: document.getElementById('itemId'),
    itemName: document.getElementById('itemName'),
    itemQty: document.getElementById('itemQty'),
    notifier: document.getElementById('notifier'),
    reloadBtn: document.getElementById('reloadBtn'),
    dismissBtn: document.getElementById('dismissBtn'),
    cancelBtn: document.getElementById('cancelBtn')
};

document.addEventListener('DOMContentLoaded', initialize);

function initialize() {
    fetchInventory();
    attachHandlers();
}

function attachHandlers() {
    elements.form.onsubmit = registerItem;
    elements.modifyForm.onsubmit = updateItem;
    elements.reloadBtn.onclick = fetchInventory;
    elements.dismissBtn.onclick = hideEditPanel;
    elements.cancelBtn.onclick = hideEditPanel;
    elements.editOverlay.onclick = (evt) => {
        if (evt.target === elements.editOverlay) hideEditPanel();
    };
}

async function fetchInventory() {
    try {
        const resp = await fetch(`${AppSettings.apiBase}${AppSettings.paths.list}`);
        if (!resp.ok) throw new Error('Network response failed');
        inventory = await resp.json();
        displayInventory();
        refreshMetrics();
    } catch (err) {
        console.error('Fetch error:', err);
        notify('Unable to load inventory data', 'error-type');
        showEmpty();
    }
}

async function registerItem(evt) {
    evt.preventDefault();
    const name = elements.nameInput.value.trim();
    const qty = parseInt(elements.qtyInput.value);
    
    if (!name || qty < 0) {
        notify('Invalid input data', 'error-type');
        return;
    }
    
    try {
        const resp = await fetch(`${AppSettings.apiBase}${AppSettings.paths.add}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ soda: name, amount: qty })
        });
        
        if (!resp.ok) throw new Error('Registration failed');
        notify(`${name} registered successfully`, 'success-type');
        elements.form.reset();
        fetchInventory();
    } catch (err) {
        console.error('Register error:', err);
        notify('Registration unsuccessful', 'error-type');
    }
}

function showEditPanel(item) {
    elements.itemId.value = item.id;
    elements.itemName.value = item.soda;
    elements.itemQty.value = item.amount;
    elements.editOverlay.classList.add('visible');
}

function hideEditPanel() {
    elements.editOverlay.classList.remove('visible');
    elements.modifyForm.reset();
}

async function updateItem(evt) {
    evt.preventDefault();
    const id = parseInt(elements.itemId.value);
    const qty = parseInt(elements.itemQty.value);
    const name = elements.itemName.value;
    
    if (qty < 0) {
        notify('Quantity cannot be negative', 'error-type');
        return;
    }
    
    try {
        const endpoint = `${AppSettings.apiBase}${AppSettings.paths.modify.replace(':id', id)}`;
        const resp = await fetch(endpoint, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ soda: name, amount: qty })
        });
        
        if (!resp.ok) throw new Error('Update failed');
        notify(`${name} updated successfully`, 'success-type');
        hideEditPanel();
        fetchInventory();
    } catch (err) {
        console.error('Update error:', err);
        notify('Update unsuccessful', 'error-type');
    }
}

async function removeItem(id, name) {
    if (!confirm(`Remove "${name}" from inventory?`)) return;
    
    try {
        const endpoint = `${AppSettings.apiBase}${AppSettings.paths.remove.replace(':id', id)}`;
        const resp = await fetch(endpoint, { method: 'DELETE' });
        
        if (!resp.ok) throw new Error('Deletion failed');
        notify(`${name} removed successfully`, 'success-type');
        fetchInventory();
    } catch (err) {
        console.error('Delete error:', err);
        notify('Deletion unsuccessful', 'error-type');
    }
}

function displayInventory() {
    if (!inventory || inventory.length === 0) {
        showEmpty();
        return;
    }
    
    elements.grid.innerHTML = inventory.map(item => {
        const safeItem = { id: item.id, soda: sanitize(item.soda), amount: item.amount };
        return `
            <div class="stock-card">
                <div class="item-id-badge">REF: ${item.id}</div>
                <div class="item-icon">${getIconFor(item.soda)}</div>
                <div class="item-title">${sanitize(item.soda)}</div>
                <div class="quantity-display">
                    <span class="quantity-value">${item.amount}</span>
                    <span class="quantity-text">in stock</span>
                </div>
                <div class="button-group">
                    <button class="action-btn modify-btn" onclick='showEditPanel(${JSON.stringify(safeItem)})'>
                        Modify
                    </button>
                    <button class="action-btn remove-btn" onclick="removeItem(${item.id}, '${sanitize(item.soda)}')">
                        Remove
                    </button>
                </div>
            </div>
        `;
    }).join('');
}

function showEmpty() {
    elements.grid.innerHTML = `
        <div class="empty-display">
            <div class="empty-symbol">📦</div>
            <h3>Inventory is empty</h3>
            <p>Start by adding your first beverage item</p>
        </div>
    `;
}

function refreshMetrics() {
    const uniqueCount = inventory.length;
    const stockTotal = inventory.reduce((acc, item) => acc + item.amount, 0);
    let topItem = '—';
    
    if (inventory.length > 0) {
        const highest = inventory.reduce((a, b) => a.amount > b.amount ? a : b);
        topItem = highest.soda;
    }
    
    document.getElementById('uniqueCount').textContent = uniqueCount;
    document.getElementById('stockTotal').textContent = stockTotal;
    document.getElementById('topItem').textContent = topItem;
}

function notify(msg, styleClass = 'success-type') {
    elements.notifier.textContent = msg;
    elements.notifier.className = `notification-bar ${styleClass} show`;
    setTimeout(() => elements.notifier.classList.remove('show'), 3500);
}

function sanitize(str) {
    const temp = document.createElement('div');
    temp.textContent = str;
    return temp.innerHTML;
}
