package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	r, err = o()
	if err != nil {
		if _, transient := err.(TransientError); transient {
			for err != nil && transient {
				r, err = o()
				_, transient = err.(TransientError)
			}

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	defer r.Close()
	defer func() {
		if rr := recover(); rr != nil {
			if f, ok := rr.(FrobError); ok {
				err = f.inner
				r.Defrob(f.defrobTag)
			} else {
				err, _ = rr.(error)
			}
		}
	}()

	r.Frob(input)

	return err
}
