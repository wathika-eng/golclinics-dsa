# import requests

# r = requests.get("https://coderbyte.com/api/challenges/json/age-counting")
# # find age >= 50
# age = r.json()["data"]
# age = age.split(", ")
# count = 0
# for i in age:
#     if "age=" in i:
#         if int(i.split("=")[1]) >= 50:
#             count += 1
# print(count)
