// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HomePage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col min-h-[100dvh]\"><header class=\"px-4 lg:px-6 h-14 flex items-center\"><a class=\"flex items-center justify-center\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-6 w-6\"><path d=\"M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4\"></path> <path d=\"M9 18c-4.51 2-5-2-7-2\"></path></svg> <span class=\"sr-only\">Jot</span></a><nav class=\"ml-auto flex gap-4 sm:gap-6\"><a class=\"text-sm font-medium hover:underline underline-offset-4\" href=\"#\">features</a> <a href=\"#\" class=\"text-sm font-medium hover:underline underline-offset-4\">Instructions</a> <a class=\"text-sm font-medium hover:underline underline-offset-4\" href=\"#\">Invite</a></nav></header><main class=\"flex-1\"><section class=\"w-full py-12 sm:py-24 md:py-32 lg:py-40 bg-[#24292f] text-white\"><div class=\"container px-4 md:px-6 flex flex-col items-center text-center space-y-6\"><div class=\"flex items-center gap-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><path d=\"M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4\"></path> <path d=\"M9 18c-4.51 2-5-2-7-2\"></path></svg></div><h1 class=\"text-3xl font-bold sm:text-4xl md:text-5xl\">GitHub Project Management Bot</h1></div><p class=\"max-w-[600px] text-lg\">A powerful GitHub-integrated Discord bot that can help you manage your projects, collaborate with your team, and more.</p></div></section><section id=\"features\" class=\"w-full py-12 sm:py-24 md:py-32 lg:py-40 bg-[#2c2f33] text-white\"><div class=\"container px-4 md:px-6 flex flex-col items-center space-y-12\"><div class=\"text-center space-y-4\"><h2 class=\"text-3xl font-bold sm:text-4xl md:text-5xl\">Key Features</h2><p class=\"max-w-[600px] text-lg\">Discover the powerful features that make our GitHub project management bot stand out.</p></div><div class=\"grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8\"><div class=\"flex flex-col items-center text-center space-y-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><path d=\"M5 7 3 5\"></path> <path d=\"M9 6V3\"></path> <path d=\"m13 7 2-2\"></path> <circle cx=\"9\" cy=\"13\" r=\"3\"></circle> <path d=\"M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17\"></path> <path d=\"M16 16h2\"></path></svg></div><h3 class=\"text-xl font-bold\">Project Management</h3><p class=\"max-w-[300px] text-lg\">Easily manage your GitHub projects and issues right from your Discord server.</p></div><div class=\"flex flex-col items-center text-center space-y-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><rect width=\"8\" height=\"8\" x=\"2\" y=\"2\" rx=\"2\"></rect> <path d=\"M14 2c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2\"></path> <path d=\"M20 2c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2\"></path> <path d=\"M10 18H5c-1.7 0-3-1.3-3-3v-1\"></path> <polyline points=\"7 21 10 18 7 15\"></polyline> <rect width=\"8\" height=\"8\" x=\"14\" y=\"14\" rx=\"2\"></rect></svg></div><h3 class=\"text-xl font-bold\">Collaboration</h3><p class=\"max-w-[300px] text-lg\">Collaborate with your team on GitHub projects directly within your Discord server.</p></div><div class=\"flex flex-col items-center text-center space-y-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><rect width=\"16\" height=\"13\" x=\"6\" y=\"4\" rx=\"2\"></rect> <path d=\"m22 7-7.1 3.78c-.57.3-1.23.3-1.8 0L6 7\"></path> <path d=\"M2 8v11c0 1.1.9 2 2 2h14\"></path></svg></div><h3 class=\"text-xl font-bold\">Notifications</h3><p class=\"max-w-[300px] text-lg\">Stay up-to-date with the latest updates on your GitHub projects right in your Discord server.</p></div></div></div></section><section id=\"instructions\" class=\"w-full py-12 sm:py-24 md:py-32 lg:py-40 bg-[#24292f] text-white\"><div class=\"container px-4 md:px-6 flex flex-col items-center space-y-12\"><div class=\"text-center space-y-4\"><h2 class=\"text-3xl font-bold sm:text-4xl md:text-5xl\">How to Add the Bot</h2><p class=\"max-w-[600px] text-lg\">Follow these simple steps to add our GitHub project management bot to your Discord server.</p></div><div class=\"grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8\"><div class=\"flex flex-col items-center text-center space-y-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><path d=\"M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71\"></path> <path d=\"M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71\"></path></svg></div><h3 class=\"text-xl font-bold\">1. Click the Invite Button</h3><p class=\"max-w-[300px] text-lg\">Click the \"Invite\" button below to start the process of adding the bot to your Discord server.</p></div><div class=\"flex flex-col items-center text-center space-y-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><rect width=\"20\" height=\"8\" x=\"2\" y=\"2\" rx=\"2\" ry=\"2\"></rect> <rect width=\"20\" height=\"8\" x=\"2\" y=\"14\" rx=\"2\" ry=\"2\"></rect> <line x1=\"6\" x2=\"6.01\" y1=\"6\" y2=\"6\"></line> <line x1=\"6\" x2=\"6.01\" y1=\"18\" y2=\"18\"></line></svg></div><h3 class=\"text-xl font-bold\">2. Select Your Server</h3><p class=\"max-w-[300px] text-lg\">Choose the Discord server you want to add the bot to.</p></div><div class=\"flex flex-col items-center text-center space-y-4\"><div class=\"bg-[#6e40c9] rounded-full p-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-10 w-10\"><path d=\"M20 6 9 17l-5-5\"></path></svg></div><h3 class=\"text-xl font-bold\">3. Authorize the Bot</h3><p class=\"max-w-[300px] text-lg\">Authorize the bot to join your server and start using its features.</p></div></div><a href=\"#\" class=\"inline-flex h-10 items-center justify-center rounded-md bg-[#6e40c9] px-8 text-sm font-medium text-white shadow transition-colors hover:bg-[#5b6eae] focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50\">Invite Bot</a></div></section></main><footer class=\"flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t\"><p class=\"text-xs text-muted-foreground\">© 2024 GitHub Project Management Bot. All rights reserved.</p><nav class=\"sm:ml-auto flex gap-4 sm:gap-6\"></nav></footer></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
