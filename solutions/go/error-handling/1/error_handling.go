package erratum

func Frobber(r Resource, input string) (err error) {
    defer func() {
        rec := recover()
        if rec != nil {
            switch e := rec.(type) {
                case FrobError:
            		r.Defrob(e.defrobTag)
                    err = e.inner
                case error:
            		err = e
            }
        }

        r.Close()
	}()

    r.Frob(input)

    return err
}

func Use(opener ResourceOpener, input string) error {
	resource, err := opener()

    switch err.(type) {
        case TransientError:
    		return Use(opener, input)
    }

    if err != nil {
        return err
    }

    err = Frobber(resource, input)

    return err
}
