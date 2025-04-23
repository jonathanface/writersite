import { PostEntry } from "../types/PostEntry";

export const NewsEntry = (props: PostEntry) => {

    return (
        <div key={props.id} className="sm:m-16 sm:max-w-[800px] news">
            <div className="flex justify-between sm:leading-6 leading-5">
                <h2 className="text-[1.5rem]">{props.title}</h2>
                <span className="text-base leading-4">{props.posted_on.toDateString()}</span>
            </div>
            <div className="sm:p-8 pb-8 pt-4 text-base font-roboto overflow-hidden" dangerouslySetInnerHTML={{ __html: props.htmlString }} />
            <hr />
        </div>
    );
}