from transformers import pipeline

text_generator = pipeline("text-generation", model="distilgpt2", truncation=True, pad_token_id=50256)
output = text_generator("", max_length=50, num_return_sequences=1)

print(output)

