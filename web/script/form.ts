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


function onChange(ev: Event) {

}

function onSubmit(ev: SubmitEvent) {
    ev.preventDefault();


    if (ev.target instanceof HTMLFormElement) {
        console.log(getValues(ev.target));
    }
}

document.addEventListener("submit", onSubmit);
document.addEventListener("change", onChange);
