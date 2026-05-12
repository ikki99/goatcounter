import re
import os

filepath = r"e:\go.cx\10w_goatcounter\i18n\zh-CN.toml"
with open(filepath, "r", encoding="utf-8") as f:
    content = f.read()

translations = {
    "button/delete-pageviews": "删除页面浏览记录",
    "button/merge": "合并",
    "button/search": "搜索",
    "button/show": "显示",
    "button/upload-ga-import": "上传 GA 导出数据",
    "dashboard/loading": "正在加载...",
    "dashboard/pages/change": "变化",
    "data-collect/help/hits": "记录所有的页面浏览和事件。",
    "data-collect/help/user-agent": "收集用户代理信息用于统计浏览器和操作系统。",
    "data-collect/label/hits": "点击量",
    "dont-lock": "不锁定",
    "email-report/daily": "每天",
    "email-report/monthly": "每月",
    "email-report/never": "从不",
    "email-report/two-weeks": "每两周",
    "email-report/weekly": "每周",
    "error/email-exists": "该邮箱已存在。",
    "error/merge-self": "不能将账号合并到它自己。",
    "error/unknown-activation-token": "无效的激活令牌或令牌已过期。",
    "header/add-api-token": "添加 API 令牌",
    "header/export-to-csv": "导出为 CSV",
    "header/fewer-number": "数字缩写",
    "header/import-from-csv": "从 CSV 导入",
    "header/import-from-ga": "从 Google Analytics 导入",
    "header/last-used-at": "最后使用时间",
    "header/manage-pageviews": "管理页面浏览记录",
    "header/merge-account": "合并账号",
    "help/allow-embed": "允许在哪些外部网站通过 iframe 嵌入仪表盘。",
    "help/configure-dashboard": "你可以自定义在仪表盘中要显示的部件以及它们的顺序。",
    "help/datepicker": "使用日期选择器来选择统计时间范围。",
    "help/email-reports": "设置定期向你的邮箱发送流量统计报告。",
    "help/fewer-numbers": "启用后大数字将使用缩写（如：将 10,000 显示为 10k）。",
    "help/fewer-numbers-lock": "为所有用户锁定大数字缩写设置。",
    "help/locked-until": "账号或功能被锁定直至该时间。",
    "help/need-cities": "需要额外配置以收集城市级别的地理位置信息。",
    "label/access-to-sites": "有权访问的站点",
    "label/campaigns": "营销活动",
    "label/dashboard-allow-embed": "允许被嵌入的网站列表",
    "label/datepicker": "日期选择",
    "label/email-reports": "定期邮件报告",
    "label/fewer-numbers": "缩写大数字",
    "label/lock-setting": "锁定此设置",
    "label/match-case": "区分大小写",
    "label/merge-account-confirmation": "确认合并账号信息",
    "label/merge-to": "合并至",
    "label/theme": "显示主题",
    "link/import": "导入数据",
    "link/manage-pageviews": "管理页面记录",
    "link/merge-account": "合并账号",
    "link/show-less": "收起",
    "nav-dash/filter-more-help": "查看更多筛选帮助",
    "nav-dash/hour": "小时",
    "nav-dash/view-by": "查看跨度：",
    "nav-fash/filter-less-help": "收起帮助",
    "new-paren": " (新)",
    "notify/import-ga-okay": "Google Analytics 数据导入完成。",
    "p/add-multiple-accounts": "你可以在这里添加多个 GoatCounter 账号。",
    "p/all-sites": "此操作将应用到你管理的所有站点。",
    "p/delete-account": "警告：这会永久删除你的账号及所有相关数据。",
    "p/export-csv": "将收集到的所有流量数据导出为 CSV 格式文件以便下载分析。",
    "p/import-ga": "支持导入 Google Analytics (GA4) 的历史数据包。",
    "p/manage-hits": "在这里你可以删除特定路径的点击数据，此操作不可逆。",
    "p/merge-account": "将当前账号的所有站点和数据无缝迁移到另一个账号中。",
    "p/num-accounts": "管理账号数量限制",
    "p/rm-hits-help": "输入你要删除的路径，系统将移除匹配的记录。",
    "theme/dark": "深色",
    "theme/light": "浅色",
    "theme/system": "跟随系统",
    "tooltip/change-period": "更改数据统计的周期",
    "top-nav/documentation": "文档",
    "validate/need-one": "必须至少提供一个有效值",
    "view not set": "未设置视图"
}

for key, value in translations.items():
    # Looking for exact match like `["key"]\n` or `["key"]\n\n`
    pattern = re.compile(r'\["?' + re.escape(key) + r'"?\]\n(?!\s*default)')
    content = pattern.sub(f'["{key}"]\n  default = "{value}"\n', content)

with open(filepath, "w", encoding="utf-8") as f:
    f.write(content)
print("Updated translations successfully.")
