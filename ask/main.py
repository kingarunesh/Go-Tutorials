a = 1



def test():
    print(a)


def main():
    global a
    test()

    print(a)
    a = 2
    print(a)


main()
