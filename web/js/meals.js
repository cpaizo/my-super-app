// 便當清單設定
const MEAL_CONFIG = [
    { name: "排骨飯", price: 120 },
    { name: "雞排飯", price: 120 },
    { name: "雞腿飯", price: 130 },
    { name: "素食便當", price: 70 },
    { name: "魚排便當", price: 130 },
    { name: "控肉飯", price: 110 } // 未來要增加直接加這一行就好
];

// 提供一個方便查詢價格的物件，不需手動修改
const mealPrices = MEAL_CONFIG.reduce((acc, meal) => {
    acc[meal.name] = meal.price;
    return acc;
}, {});