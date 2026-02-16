const SUPABASE_URL = "https://efkiuxgwlzjftlqxtett.supabase.co";
const SUPABASE_KEY = "sb_publishable_G2itMMaOhbauxhKrnUJtwA_XDbhELbU";

// 這裡加上 window. 前綴，確保全域都能抓到這個物件
window._supabase = supabase.createClient(SUPABASE_URL, SUPABASE_KEY);