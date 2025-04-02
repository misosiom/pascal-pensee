import json
import os
from openai import OpenAI
from typing import List
from dotenv import load_dotenv

# 🔑 .envファイルから環境変数を読み込む
load_dotenv()

# OpenAIクライアントの初期化（envから取得）
client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))

# 英語のpenseesデータをID指定で読み込む
def load_passages(filepath: str, id_range: List[int]) -> List[dict]:
    with open(filepath, "r", encoding="utf-8") as f:
        data = json.load(f)
    return [item for item in data if item["id"] in id_range]

# GPTで翻訳
def translate_text(text: str) -> str:
    prompt = f"次の英語の文章を丁寧で自然な日本語に翻訳してください：\n\n{text}"
    response = client.chat.completions.create(
        model="gpt-4",
        messages=[{"role": "user", "content": prompt}],
        temperature=0.7
    )
    return response.choices[0].message.content.strip()

# 翻訳と保存
def translate_and_save(input_path: str, output_path: str, ids: List[int]):
    passages = load_passages(input_path, ids)
    translated = []

    for item in passages:
        print(f"Translating ID {item['id']}...")
        jp_text = translate_text(item["text"])
        translated.append({
            "id": item["id"],
            "text": jp_text
        })

    with open(output_path, "w", encoding="utf-8") as f:
        json.dump(translated, f, ensure_ascii=False, indent=2)
    print(f"✅ 翻訳結果を保存しました: {output_path}")

# 実行部
if __name__ == "__main__":
    INPUT_JSON = "pensees.json"
    OUTPUT_JSON = "pensees_ja_1_20.json"
    ID_RANGE = list(range(1, 923))

    translate_and_save(INPUT_JSON, OUTPUT_JSON, ID_RANGE)
