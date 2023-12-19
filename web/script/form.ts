type FieldValues = Record<string, string>;

function registerForms() {
    document.addEventListener("change", onChange);
    document.addEventListener("submit", onSubmit);

    function onSubmit(ev: SubmitEvent) {
        ev.preventDefault();
    }

    function onChange(ev: Event) {
        if (ev.target instanceof HTMLInputElement) {
            onChangeInput(ev.target);
        }
    }

    function getErrorFlag(validity: ValidityStateFlags): boolean {
        return Boolean(
            validity.valueMissing
            || validity.patternMismatch
        );
    }

    function getErrorMessage(validity: ValidityStateFlags): string {
        switch (true) {
            case validity.valueMissing:
                return "Field is required";

            case validity.patternMismatch:
                return "Invalid format";

            default:
                return "";
        }
    }

    function onChangeInput({ form, name, validity }: HTMLInputElement) {
        if (!form) {
            return;
        }

        const errorElement = form.elements.namedItem(`${name}-error`);

        if (errorElement instanceof HTMLOutputElement) {
            errorElement.hidden = !getErrorFlag(validity);
            errorElement.textContent = getErrorMessage(validity);
        }
    }
}

registerForms();
