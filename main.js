const longestPrefix = (strings) => {
    let prefix = 'empty';

    if (!strings.length) {
        return prefix;
    }

    prefix = '';

    let flag = true

    for (let i = 0; i < strings[0].length; i++) {
        let letter = strings[0].charAt(i);

        for (let j = 1; j < strings.length; j++) {
            if (strings[j].charAt(i) !== letter) {
                letter = null;
            }
        }

        if (letter) {
            prefix += letter;
        } else {
            break;
        }
    }

    return prefix;
}

console.log(
    longestPrefix(
        [
            "flower",
            "flow",
            "flight",
            "flon"
        ]
    )
)