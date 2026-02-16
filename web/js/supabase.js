// web/js/supabase.js
const SUPABASE_URL = "https://efkiuxgwlzjftlqxtett.supabase.co";
const SUPABASE_KEY = "sb_publishable_G2itMMa0hbauxhKrnUJtwA_XDbhELbU";

// 使用 window._supabase 確保在所有頁面都能正確存取
window._supabase = supabase.createClient(SUPABASE_URL, SUPABASE_KEY);