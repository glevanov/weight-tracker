import type { Locale } from "../i18n";

export const zhTW: Locale = {
  name: "繁體中文",
  login: {
    login: "登入",
    password: "密碼",
    submit: "提交",
    emptyCredentials: "使用者名稱和密碼不能為空",
  },
  addWeight: {
    header: "輸入你的體重",
    success: "體重已成功添加",
    submit: "提交",
    showGraph: "顯示圖表",
    empty: "重量不能為空",
  },
  chart: {
    twoWeeks: "兩週",
    month: "一個月",
    quarter: "季度",
    year: "一年",
    allData: "所有資料",
    addWeight: "添加體重",
    loading: "載入中！",
    errorOccurred: "發生錯誤",
  },
  initialLoading: {
    loading: "正在啟動伺服器！這可能需要幾分鐘。",
    failed: "啟動伺服器失敗。請嘗試刷新頁面或稍後再試。",
  },
};
