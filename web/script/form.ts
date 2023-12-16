type FieldValues = Record<string, string>;

function getValues(form: HTMLFormElement): FieldValues {
    const values: FieldValues = {};

    for (const element of form.elements) {
        if (element instanceof HTMLInputElement) {
            values[element.name] = element.value;
        }
    }

    return values;
}

class Form {
    private form: Maybe<HTMLFormElement> = null;

    public constructor(name: string) {
        this.form = document.getElementById(name) as Maybe<HTMLFormElement>;

        if (this.form) {
            this.form.addEventListener("blur", this.onBlur);
            this.form.addEventListener("submit", this.onSubmit);
        }
    }

    private onSubmit = (ev: SubmitEvent) => {
        ev.preventDefault();

        if (ev.target instanceof HTMLFormElement) {
            console.log(getValues(ev.target));
        }
    }

    private onBlur = (ev: Event) => {
        if (ev.target instanceof HTMLInputElement && ev.target.form) {
            console.log(getValues(ev.target.form));
        }
    }
}
