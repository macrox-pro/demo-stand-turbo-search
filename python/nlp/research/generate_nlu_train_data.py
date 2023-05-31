import yaml
import random
import pymorphy2

from string import Template

morph = pymorphy2.MorphAnalyzer()


def inflect_word(word, grams):
    forms = morph.parse(word)
    if len(forms) < 1:
        return word

    form = forms[0].inflect(grams)
    if form is None:
        return word

    return form.word


def inflect(text, grams):
    if len(grams) < 1:
        return text

    words = text.split()
    for index, word in enumerate(words):
        words[index] = inflect_word(word, grams)

    return " ".join(words)


def modificate(text, mod):
    if mod == "lower":
        return text.lower()
    elif mod == "upper":
        return text.upper()
    elif mod == "title":
        return text.title()

    return text


def y2l(tol):
    if type(tol) == str:
        tol = tol.split('\n')
        for ii, iv in enumerate(tol):
            tol[ii] = iv.strip().lstrip('-').strip()

    return tol


if __name__ == '__main__':
    with open("../../../.index.counts.yaml", "r") as stream:
        data = yaml.safe_load(stream)

    with open("../data/train_templates.yml", "r") as stream:
        instructions = yaml.safe_load(stream)

    types = data.get('types')
    for index, value in enumerate(types):
        types[index] = value.get('name')

    genres = data.get('genres')
    for index, value in enumerate(genres):
        genres[index] = value.get('name')

    persons = data.get('persons')
    for index, value in enumerate(persons):
        persons[index] = value.get('name')

    print(f'version: \"{instructions.get("version")}"\n')
    print('nlu:')
    for nlu in instructions.get('nlu', []):
        if "regex" in nlu:
            print(f'\n- regex: {nlu.get("regex")}')
        elif "lookup" in nlu:
            print(f'\n- lookup: {nlu.get("lookup")}')
        elif "intent" in nlu:
            print(f'\n- intent: {nlu.get("intent")}')
        else:
            continue

        print('  examples: |')

        for example in y2l(nlu.get('examples', [])):
            if len(example) > 0:
                print(f'    - {example}')

        mods = y2l(nlu.get('mods', []))
        entries = nlu.get('entries', dict())

        for value in nlu.get('templates', []):

            template = Template(value)

            if nlu.get("lookup") == "genre":
                results = []

                genres_entry = entries.get("genres", dict())

                for genre in genres_entry.get("items", genres):

                    for mod in mods:

                        if len(genres_entry.get("grams", {})) > 0:

                            for grams in genres_entry.get("grams", {}):
                                results.append(template.substitute(
                                    genre=modificate(inflect(genre, set(grams)), mod),
                                ).strip())
                        else:

                            results.append(template.substitute(
                                genre=modificate(genre, mod),
                            ).strip())

                if len(results) > 0:
                    for result in list(set(results)):
                        print(f'    - {result.strip()}')

            elif nlu.get("intent", "").find("person") >= 0:

                for i in range(min(121, len(persons))):

                    person = random.choice(persons).split()

                    if bool(random.getrandbits(1)) and len(person) > 1:
                        person = person[len(person) - 1]
                    else:
                        person = ' '.join(person)

                    mod = random.choice(mods)

                    types_entry = entries.get("types", dict())
                    genres_entry = entries.get("genres", dict())
                    persons_entry = entries.get("persons", dict())

                    t = random.choice(types_entry.get("items", {}))
                    if len(types_entry.get("grams", [])) > 0:
                        t = inflect(t, set(random.choice(types_entry.get("grams", {}))))

                    genres = genres_entry.get("items", genres)
                    if len(genres) > 0:
                        genre = random.choice(genres)
                        if len(genres_entry.get("grams", {})) > 0:
                            genre = inflect(genre, set(random.choice(genres_entry.get("grams", {}))))
                    else:
                        genre = ""

                    if len(persons_entry.get("grams", {})) > 0:
                        person = inflect(person, set(random.choice(persons_entry.get("grams", {}))))

                    if template.template.find("(details)") >= 0:
                        details = random.choice(y2l(entries.get('details', {})))
                        if len(details) < 1:
                            continue

                        result = template.substitute(
                            type=modificate(t, mod),
                            genre=modificate(genre, mod),
                            person=modificate(person, mod),
                            details=details
                        )
                    else:
                        result = template.substitute(
                            type=modificate(t, mod),
                            genre=modificate(genre, mod),
                            person=modificate(person, mod),
                        )

                    if len(result.strip()) > 0:
                        print(f'    - {result.strip()}')
