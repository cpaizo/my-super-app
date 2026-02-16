// web/js/supabase.js
const SUPABASE_URL = "https://efkiuxgwlzjftlqxtett.supabase.co";

// 請確保這裡貼上的是 Supabase 後台面板中 Project API 下方的 anon key
const SUPABASE_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."; 

window._supabase = supabase.createClient(SUPABASE_URL, SUPABASE_KEY);