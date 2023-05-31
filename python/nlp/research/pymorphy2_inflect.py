import pymorphy2

morph = pymorphy2.MorphAnalyzer()


def inflect_word(word, grams):
    forms = morph.parse(word)
    if len(forms) < 1:
        return word

    form = forms[0].inflect(grams)
    if form is None:
        return word

    return form.word


def make_agree_with_number(word, num):
    forms = morph.parse(word)
    if len(forms) < 1:
        return word

    form = forms[0].make_agree_with_number(num)
    if form is None:
        return word

    return form.word


if __name__ == '__main__':
    for word in {"фильм", "викторина", "Александр", "Иванов"}:
        print(f'--- {word} ---')
        for gram in {'nomn', 'gent', 'datv', 'accs', 'ablt', 'loct', 'voct', 'gen2', 'acc2', 'loc2'}:
            print(gram, 'sign', inflect_word(word, {gram, 'sing'}))
            print(gram, 'plur', inflect_word(word, {gram, 'plur'}))

