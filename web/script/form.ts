type FieldValues = Record<string, string>;

function registerForms() {
    document.addEventListener("change", onChange);
    document.addEventListener("submit", onSubmit);

    function onSubmit(ev: SubmitEvent) {
        ev.preventDefault();

        if (ev.target instanceof HTMLFormElement) {
            const isValid = Array.from(ev.target.elements).every((el) => {
                if (el instanceof HTMLInputElement) {
                    return isInputValid(el)
                }

                return false;
            });

            if (isValid) {
                console.log("Submitting");
            } else {
                console.log("Form is invalid");
            }
        }
    }

    function onChange(ev: Event) {
        if (ev.target instanceof HTMLInputElement) {
            onChangeInput(ev.target);
        }
    }

    function isInputValid({ validity }: HTMLInputElement): boolean {
        return (
            validity.badInput
            || validity.customError
            || validity.patternMismatch
            || validity.rangeOverflow
            || validity.rangeUnderflow
            || validity.stepMismatch
            || validity.tooLong
            || validity.tooShort
            || validity.typeMismatch
            || validity.valid
            || validity.valueMissing
        );
    }

    function onChangeInput(input: HTMLInputElement) {
        if (!input.form || isInputValid(input)) {
            return;
        }

        const error = input.form.elements.namedItem(`${input.name}-error`);

        if (error instanceof HTMLOutputElement) {
            error.textContent = getErrorMessage(input);
            error.hidden = error.textContent === null;
        }
    }

    function getErrorMessage({ validity }: HTMLInputElement): string | null {
        switch (true) {
            case validity.badInput:
                return "badInput";

            case validity.customError:
                return "customError";

            case validity.patternMismatch:
                return "patternMismatch";

            case validity.rangeOverflow:
                return "rangeOverflow";

            case validity.rangeUnderflow:
                return "rangeUnderflow";

            case validity.stepMismatch:
                return "stepMismatch";

            case validity.tooLong:
                return "tooLong";

            case validity.tooShort:
                return "tooShort";

            case validity.typeMismatch:
                return "typeMismatch";

            case validity.valueMissing:
                return "valueMissing";

            default:
                return null;
        }
    }
}

registerForms();
