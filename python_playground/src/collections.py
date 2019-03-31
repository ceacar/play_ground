# Complete the hackerCards function below.
def hackerCards(collection, d):
    counting_sort_arr = [0] * (10**9)
    for card_value in collection:
        counting_sort_arr[card_value] = 1


    cards_to_buy = []
    total_face_value = 0
    for idx in range(10**9):
        if counting_sort_arr[idx] == 0:
            total_face_value += idx
            if total_face_value > d:
                break
            else:
                cards_to_buy.append(idx)
    return cards_to_buy



if __name__ == '__main__':
    print(hackerCards([1,3,4], 7))
