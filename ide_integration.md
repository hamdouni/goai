# 2024-02-22

I used Gen to integrate ollama with neovim.

I used this plugin configuration :

```lua
return {
	{
		"David-Kunz/gen.nvim",
		opts = {
			model = "deepseek-coder:6.7b", -- The default model to use.
			display_mode = "split", -- The display mode. Can be "float" or "split".
			show_prompt = false, -- Shows the Prompt submitted to Ollama.
			show_model = false, -- Displays which model you are using at the beginning of your chat session.
			no_auto_close = false, -- Never closes the window automatically.
			init = function(options) pcall(io.popen, "ollama serve > /dev/null 2>&1 &") end,
			-- Function to initialize Ollama
			command = "curl --silent --no-buffer -X POST http://localhost:11434/api/generate -d $body",
			-- The command for the Ollama service. You can use placeholders $prompt, $model and $body (shellescaped).
			-- This can also be a lua function returning a command string, with options as the input parameter.
			-- The executed command must return a JSON object with { response, context }
			-- (context property is optional).
			list_models = '<omitted lua function>', -- Retrieves a list of model names
			debug = false -- Prints errors and the command which is run.
		}
	},
}
```

Now, I can select some code then use :Gen Review_Code

## Using Gen.nvim and ollama with model `deepseek-coder:6.7b`

For example, I selected this Go code :

```go
func SpaceRightPad(value string, length int) (string, error) {
	if len(value) > length {
		return "", fmt.Errorf("value %s to format is too big %d", value, length)
	}
	return value + strings.Repeat(" ", length-len(value)), nil
}
```

and asked for his review and got this :

    ---
    The function `SpaceRightPad` pads the input string with spaces on its right side up to a specified length. However, this implementation is inefficient as it repeatedly calculates and appends the same number of whitespace characters unnecessarily. The unnecessary repetition can be minimized by only appending one instance of all required padding:

    ```go
    func SpaceRightPad(value string, length int) (string, error) {
        if len(value) > length {
            return "", fmt.Errorf("value %s to format is too big %d", value, length)
        }
        return value + strings.Repeat(" ", length-len(value)), nil
    }
    ```
    This version of `SpaceRightPad` simply calculates how many spaces are required and append them once at the end, which makes it more efficient for large inputs.
    The error checking part is good as well; if the input value's length exceeds the provided padding length then an appropriate error message will be returned.

It gives me the exact same function after complaining about some unnecessary repetition that does not exist.

## Change model to `mistral`

So, i tried with `mistral` model :

    The `SpaceRightPad` function takes a string `value` and an integer `length`, and returns a padded string and an error if any. Here are some suggestions:

    1. The function name could be more descriptive, such as `RightPadWithSpaces`.
    2. Check for a negative length argument and return an error accordingly.
    3. Use constants or named variables instead of magic numbers (`len(value)` vs. `valueLen`, `length-len(value)` vs. `paddingLength`).
    4. Consider adding documentation to explain the function's purpose and behavior.
    5. Use Go's built-in `bytes.Repeat` function instead of `strings.Repeat` if the output is expected to be bytes, not strings.

    Here is an example of how the code could be refactored based on these suggestions:

    ```go
    // RightPadWithSpaces pads a given string with spaces on the right side up to a given length.
    func RightPadWithSpaces(value string, length int) (string, error) {
        if length < 0 {
            return "", fmt.Errorf("invalid length: %d", length)
        }

        valueLen := len(value)
        if valueLen > length {
            return "", fmt.Errorf("value %s to format is too big %d", value, length)
        }

        padding := bytes.Repeat([]byte(" "), length-valueLen)
        paddedValue := value + padding

        return paddedValue, nil
    }
    ```

It seems more accurate than the deepcode version.

The second point suggesting to check for negative length would be ok in other contest,
but here, this test is implicitly included in the comparaison between the value
length and the padding length.

The third point about using a variable instead of len(value) is legit, also
because I use it elsewhere.

The fourth point is not important : I didn't select the comment submited to the
AI.

And the last point doesn't make sense in this context : I explicitly return a string and not bytes.

