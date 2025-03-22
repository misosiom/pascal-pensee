from bs4 import BeautifulSoup
import json

# 入力HTMLファイルパス
html_file = "pg18269-images.html"

# 出力JSONファイルパス
json_file = "pensees.json"

# ファイル読み込み
with open(html_file, "r", encoding="utf-8") as f:
    soup = BeautifulSoup(f, "html.parser")

# 抽出結果
pensees = []

# すべての h4 タグを探す
for h4 in soup.find_all("h4"):
    a_tag = h4.find("a")
    if a_tag and a_tag.get("id", "").startswith("p_"):
        pensee_id = a_tag["id"].replace("p_", "")
        
        text = ""
        next = h4.find_next_sibling()
        # print(next.name)
        while next is not None and next.name == "p":
            text = text + next.get_text(strip=True)
            # print(next.name)
            next = next.find_next_sibling()

        pensees.append({
            "id": int(pensee_id),
            "text": text
        })

# JSONに保存
with open(json_file, "w",   encoding="utf-8") as f:
    json.dump(pensees, f, indent=2, ensure_ascii=False)

print(f"{len(pensees)} 個の名言を抽出し、'{json_file}' に保存しました。")

