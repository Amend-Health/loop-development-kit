namespace OliveHelpsLDK.Whispers.List
{
    /// <summary>
    /// Base class for all List elements.
    /// </summary>
    public abstract class ListBase
    {
        /// <summary>
        /// The field order. Must be at least zero. List entries are ordered by order, then by key. 
        /// </summary>
        public uint Order { get; set; }
        /// <summary>
        /// Whether the field is visible when the whisper is expanded or not.
        /// TODO: Confirm this is accurate.
        /// </summary>
        public bool Extra { get; set; }
    }

    /// <summary>
    /// The style to apply to the entry.
    /// </summary>
    public enum Style
    {
        /// <summary>
        /// Standard styling.
        /// </summary>
        None,

        /// <summary>
        /// TODO: Find this out.
        /// </summary>
        Success,

        /// <summary>
        /// TODO: Find this out.
        /// </summary>
        Warn,

        /// <summary>
        /// TODO: Find this out.
        /// </summary>
        Error
    }

    /// <summary>
    /// Defines how the entry is aligned.
    /// </summary>
    public enum Align
    {
        /// <summary>
        /// Left aligned.
        /// </summary>
        Left,

        /// <summary>
        /// Center aligned.
        /// </summary>
        Center,

        /// <summary>
        /// Right aligned.
        /// </summary>
        Right
    }

    /// <summary>
    /// Defines a label/value pair.
    /// </summary>
    public class ListPair : ListBase
    {
        /// <summary>
        /// The Label presented.
        /// </summary>
        public string Label { get; set; }

        /// <summary>
        /// The value presented as string.
        /// </summary>
        public string Value { get; set; }

        /// <summary>
        /// The entry's style.
        /// </summary>
        public Style Style { get; set; }

        /// <summary>
        /// Whether the field can be copied or not.
        /// TODO: Confirm
        /// </summary>
        public bool Copyable { get; set; }
    }

    /// <summary>
    /// A message entry.
    /// </summary>
    public class ListMessage : ListBase
    {
        /// <summary>
        /// The header of the message.
        /// </summary>
        public string Header { get; set; }

        /// <summary>
        /// The body of the message.
        /// </summary>
        public string Body { get; set; }

        /// <summary>
        /// The style the message is presented in.
        /// </summary>
        public Style Style { get; set; }

        /// <summary>
        /// The alignment of the message.
        /// </summary>
        public Align Align { get; set; }
    }

    /// <summary>
    /// A divider entry.
    /// </summary>
    public class ListDivider : ListBase
    {
        /// <summary>
        /// The style of the divider.
        /// </summary>
        public Style Style { get; set; }
    }

    public class ListLink : ListBase
    {
        public Align Align { get; set; }
        public string Href { get; set; }
        public Style Style { get; set; }
        public string Text { get; set; }
    }
}