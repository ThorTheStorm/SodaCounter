const AppSettings = {
    apiBase: 'http://localhost:8080',
    paths: {
        list: '/soda',
        single: '/soda/:id',
        add: '/soda',
        modify: '/soda/:id',
        remove: '/soda/:id'
    }
};

const BeverageIcons = {
    'coca-cola': '🥤', 'coke': '🥤', 'pepsi': '🥤',
    'sprite': '💚', 'fanta': '🟠', 'dr pepper': '🟤',
    'mountain dew': '💛', '7up': '💚', 'root beer': '🟤',
    'orange': '🍊', 'grape': '🍇', 'lemon': '🍋',
    'ginger': '💛', 'cola': '🥤', 'fallback': '🥤'
};

function getIconFor(name) {
    const normalized = name.toLowerCase();
    for (let [key, icon] of Object.entries(BeverageIcons)) {
        if (normalized.includes(key)) return icon;
    }
    return BeverageIcons.fallback;
}
